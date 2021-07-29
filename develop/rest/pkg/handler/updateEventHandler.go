package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Error method")
	}
	http.ServeFile(w, r, "hello.html")
}
