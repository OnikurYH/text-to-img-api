package test

import (
	"encoding/json"
	"image/png"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"onikur.com/text-to-img-api/conf"

	"onikur.com/text-to-img-api/api"
)

func TestListFonts(t *testing.T) {
	conf.Init()
	handler := &api.ListFontsHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode >= 400 {
		t.Fatalf("Received invalid response: %d\n", res.StatusCode)
	}
	var fonts []string
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&fonts); err != nil {
		t.Fatal(err)
	}
}

func TestMakeText(t *testing.T) {
	conf.Init()
	handler := &api.MakeTextHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	res, err := http.Get(server.URL + "/" + url.PathEscape("Hello, World!") + "?fcolor=FFAA00FF")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode >= 400 {
		t.Fatalf("Received invalid response: %d\n", res.StatusCode)
	}
	defer res.Body.Close()

	img, err := png.Decode(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if img.Bounds().Size().X <= 1 || img.Bounds().Size().Y <= 1 {
		t.Fatalf("Image invalid size: (%d, %d)\n", img.Bounds().Size().X, img.Bounds().Size().Y)
	}
}
