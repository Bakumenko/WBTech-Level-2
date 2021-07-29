package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DEL" {
		fmt.Println("Error method")
	}
	http.ServeFile(w, r, "hello.html")
}
