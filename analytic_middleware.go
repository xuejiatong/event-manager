package main

import (
	"net/http"
	"strings"
)

func ReferrerTracker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Referer"]) == 1 {
			var sources = [...]string{"http://som.yale.edu",
				"http://divinity.yale.edu",
				"http://medicine.yale.edu",
				"http://law.yale.edu",
				"http://search.yale.edu"}

			referrer := r.Header["Referer"][0]
			for _, source := range sources {
				if strings.HasPrefix(referrer, source) {
					addReferrer(referrer)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func SupportTracker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/donate") {
			addSupport("donate")
		} else if strings.HasSuffix(r.URL.Path, "/support") {
			addSupport("support")
		}

		next.ServeHTTP(w, r)
	})
}
