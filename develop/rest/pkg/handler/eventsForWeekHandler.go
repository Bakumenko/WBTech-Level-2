package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Error method")
	}
	http.ServeFile(w, r, "hello.html")
}
