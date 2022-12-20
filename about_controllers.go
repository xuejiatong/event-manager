package main

import (
	"net/http"
)

func aboutController(w http.ResponseWriter, r *http.Request) {
	tmpl["about"].Execute(w, nil)
}
