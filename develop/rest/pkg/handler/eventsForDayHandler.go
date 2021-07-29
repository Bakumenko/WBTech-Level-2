package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Error method")
		return
	}

	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	events, err := h.repos.GetEventsForDay(today, tomorrow)
	if err != nil {
		return
	}

	eventsBytes, _ := json.MarshalIndent(events, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(eventsBytes)
}
