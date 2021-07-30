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
	err = handleRequests(router, hndlr)
	if err != nil {
		return err
	}

	return nil
}

func handleRequests(router *http.ServeMux, h *handler.Handler) error {
	http.HandleFunc("/events_for_day", h.GetEventsForDay)
	http.HandleFunc("/events_for_week", h.GetEventsForWeek)
	http.HandleFunc("/events_for_month", h.GetEventsForMonth)
	http.HandleFunc("/create_event", h.CreateEvent)
	http.HandleFunc("/update_event", h.UpdateEvent)
	http.HandleFunc("/delete_event", h.DeleteEvent)
	//configuredRouter := ExampleMiddleware(router)
	return http.ListenAndServe(":"+webPort, nil)
}

func ExampleMiddleware(h http.Handler) http.Handler {
	// We wrap our anonymous function, and cast it to a http.HandlerFunc
	// Because our function signature matches ServeHTTP(w, r), this allows
	// our function (type) to implicitly satisify the http.Handler interface.
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(rw, r) // serve the original request

		duration := time.Since(start)

		// log request details
		logrus.WithFields(logrus.Fields{
			"uri":      uri,
			"method":   method,
			"duration": duration,
		})
	}
	return http.HandlerFunc(logFn)
}
