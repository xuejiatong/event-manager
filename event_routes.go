package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type eventDetailContextData struct {
	Event            Event
	Errors           string
	ConfirmationCode string
}

func CreateEventPage(w http.ResponseWriter, r *http.Request) {
	tmpl["createEvent"].Execute(w, nil)
}

func CreateEvent(w http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	location := req.FormValue("location")
	image := req.FormValue("image")
	date := req.FormValue("date")

	datetime, error := time.Parse("2006-01-02T15:04", date)
	if error != nil {
		tmpl["createEvent"].Execute(w, nil)
		return
	}

	e := Event{
		Title:    title,
		Location: location,
		Image:    image,
		Date:     datetime,
	}

	_, err := addEvent(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl["createEvent"].Execute(w, nil)
}

func GetEvent(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	e, err := getEventByID(id)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	a, err := getAttendees(e)
	if err != nil {
		http.Error(w, "Attendees not found", http.StatusInternalServerError)
		return
	}

	e.Attending = a

	contextData := eventDetailContextData{
		Event:            e,
		ConfirmationCode: "",
		Errors:           "",
	}
	tmpl["event"].Execute(w, contextData)
}

func UpdateEvent(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	e, err := getEventByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// determine confirmation code
	errors := ""
	confirmationCode := ""
	email := req.FormValue("name")
	// raise error if the email is not valid
	if canRSVP(email) {
		_, err = addAttendee(e, email)
		confirmationCode = confirmationCodeForEmail(email)
	} else {
		errors = "Error: please register with a valid Yale email!"
	}

	contextData := eventDetailContextData{
		Event:            e,
		ConfirmationCode: confirmationCode,
		Errors:           errors,
	}
	if err != nil {
		tmpl["event"].Execute(w, contextData)
		return
	}

	attendees, err := getAttendees(e)
	if err != nil {
		tmpl["event"].Execute(w, contextData)
		return
	}
	// successfully add, update event
	contextData.Event.Attending = attendees

	tmpl["event"].Execute(w, contextData)
}

func EventRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/new", CreateEventPage)
	r.Post("/new", CreateEvent)
	r.Get("/{id}", GetEvent)
	r.Post("/{id}", UpdateEvent)
	return r
}
