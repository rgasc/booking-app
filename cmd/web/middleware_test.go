package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var handler testHandler
	h := NoSurf(&handler)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var handler testHandler
	h := SessionLoad(&handler)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
