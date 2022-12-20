package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Analytic struct {
	Category   string
	Count      int
	Percentage float32
}

type AnalyticContext struct {
	Analytics []Analytic
}

func ViewReferrers(w http.ResponseWriter, r *http.Request) {
	referrers, err := getReferrers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make(map[string]int)
	for _, referrer := range referrers {
		result[referrer.Referrer]++
	}

	total := 0
	for _, v := range result {
		total += v
	}

	analytics := make([]Analytic, 0)
	for k, v := range result {
		analytics = append(analytics, Analytic{k, v, float32(v) * 100 / float32(total)})
	}

	context := AnalyticContext{analytics}
	tmpl["analytics"].Execute(w, context)
}

func ViewSupport(w http.ResponseWriter, r *http.Request) {
	support, err := getSupport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make(map[string]int)
	for _, sup := range support {
		result[sup.Support]++
	}

	total := 0
	for _, v := range result {
		total += v
	}

	analytics := make([]Analytic, 0)
	for k, v := range result {
		analytics = append(analytics, Analytic{k, v, float32(v) * 100 / float32(total)})
	}

	context := AnalyticContext{analytics}
	tmpl["analytics"].Execute(w, context)
}

func AnalyticRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/referrers", ViewReferrers)
	r.Get("/supports", ViewSupport)
	return r
}
