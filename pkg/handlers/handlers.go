package handlers

import (
	"net/http"

	"github.com/rgasc/booking-app/pkg/config"
	"github.com/rgasc/booking-app/pkg/models"
	"github.com/rgasc/booking-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// stringMap := map[string]string{}
	// stringMap["test"] = "this is a test"

	// remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	// stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

// RoomOne is a room page handler
func (m *Repository) RoomOne(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "room-1.page.tmpl", &models.TemplateData{})
}

// RoomTwo is a room page handler
func (m *Repository) RoomTwo(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "room-2.page.tmpl", &models.TemplateData{})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
