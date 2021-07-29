package main

import (
	"fmt"
	"net/http"
	"simplerest/pkg/handler"
	"simplerest/pkg/repository"
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
	err = handleRequests(hndlr)
	if err != nil {
		return err
	}

	return nil
}

func handleRequests(h *handler.Handler) error {
	http.HandleFunc("/events_for_day", h.GetEventsForDay)
	http.HandleFunc("/events_for_week", h.GetEventsForWeek)
	http.HandleFunc("/events_for_month", h.GetEventsForMonth)
	http.HandleFunc("/create_event", h.CreateEvent)
	http.HandleFunc("/update_event", h.UpdateEvent)
	http.HandleFunc("/delete_event", h.DeleteEvent)
	return http.ListenAndServe(":"+webPort, nil)
}
