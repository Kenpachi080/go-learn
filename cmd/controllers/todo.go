package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	str := fmt.Sprintf("Hello world %d", id)
	w.Write([]byte(str))

	return
}
