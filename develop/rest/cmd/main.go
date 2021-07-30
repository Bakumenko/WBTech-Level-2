package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"simplerest/pkg/handler"
	"simplerest/pkg/repository"
	"time"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	db, err := repository.NewPostgresDB()
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	hndlr := handler.NewHandler(repos)
	router := http.NewServeMux()

	addr := "127.0.0.1:" + webPort
	logrus.WithField("addr", addr).Info("starting server")

	err = startServer(router, hndlr)
	if err != nil {
		return err
	}

	return nil
}

func startServer(router *http.ServeMux, h *handler.Handler) error {
	router.HandleFunc("/events_for_day", h.GetEventsForDay)
	router.HandleFunc("/events_for_week", h.GetEventsForWeek)
	router.HandleFunc("/events_for_month", h.GetEventsForMonth)
	router.HandleFunc("/create_event", h.CreateEvent)
	router.HandleFunc("/update_event", h.UpdateEvent)
	router.HandleFunc("/delete_event", h.DeleteEvent)
	configuredRouter := ExampleMiddleware(router)
	return http.ListenAndServe(":"+webPort, configuredRouter)
}

func ExampleMiddleware(h http.Handler) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		duration := time.Since(start)

		logrus.WithFields(logrus.Fields{
			"uri":      uri,
			"method":   method,
			"duration": duration,
		}).Info("input request")

		h.ServeHTTP(rw, r)
	}
	return http.HandlerFunc(logFn)
}
