package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stasxgaming/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var a *config.AppConfig

	mux := routes(a)

	switch mux.(type) {
	case *chi.Mux:
	default:
		t.Error("routes() return is not a type of mux")
	}
}
