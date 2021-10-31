package main

import (
	"net/http"
	"testing"
)

func TestCreateCSRFToken(t *testing.T) {
	var myH myHandler
	h := CreateCSRFToken(&myH)

	switch h.(type) {
	case http.Handler:
	default:
		t.Error("not http.Handler in return")
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch h.(type) {
	case http.Handler:
	default:
		t.Error("not http.Handler in return")
	}
}
