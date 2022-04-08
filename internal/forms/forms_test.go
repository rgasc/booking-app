package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	if !form.Valid() {
		t.Error("got invalid when it should have been valid")
	}
}

func TestRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)
	form.Required("1", "2", "3")

	if form.Valid() {
		t.Error("valid when it should have been invalid")
	}

	postData := url.Values{}
	postData.Add("1", "1")
	postData.Add("2", "2")
	postData.Add("3", "3")

	r = httptest.NewRequest("POST", "/", nil)
	form = New(postData)
	form.Required("1", "2", "3")

	if !form.Valid() {
		t.Error("got invalid when it should have been valid")
	}
}

func TestHas(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	if form.Has("test") {
		t.Error("found the test field when it should not have")
	}

	postData = url.Values{}
	postData.Add("test", "test")
	form = New(postData)

	if !form.Has("test") {
		t.Error("did not find the test field when it should have")
	}
}

func TestMinLength(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.MinLength("test", 5)
	if form.Valid() {
		t.Error("forms shows min length for nonexistent field")
	}

	err := form.Errors.Get("test")
	if err == "" {
		t.Error("should have an error but did not get one")
	}

	postData = url.Values{}
	postData.Add("test", "test")
	form = New(postData)

	form.MinLength("test", 100)
	if form.Valid() {
		t.Error("gives valid length when it is invalid")
	}

	postData = url.Values{}
	postData.Add("test", "test")
	form = New(postData)

	form.MinLength("test", 4)
	if !form.Valid() {
		t.Error("gives invalid length when it is valid")
	}

	err = form.Errors.Get("test")
	if err != "" {
		t.Error("should not have an error but got one")
	}
}

func TestIsEmail(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.IsEmail("test")

	if form.Valid() {
		t.Error("forms shows valid email for nonexistent field")
	}

	postData.Add("test", "test@test.com")
	form = New(postData)

	form.IsEmail("test")
	if !form.Valid() {
		t.Error("got an invalid email when it is valid")
	}

	postData = url.Values{}
	postData.Add("test", "test@test")
	form = New(postData)

	form.IsEmail("test")
	if form.Valid() {
		t.Error("got a valid email when it is invalid")
	}
}
