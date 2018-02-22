package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		result Server
		err    error
	}{
		{result: Server{}, err: nil},
	}
	for _, test := range tests {
		result, err := newServer("localhost")
		assert.IsType(t, test.err, err)
		assert.IsType(t, &test.result, result)
	}
}

func (s *Server) TestSearch(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8000/search", nil)
	if err != nil {
		t.Fatalf("could not create a request %v", err)
	}

	w := httptest.NewRecorder()

	s.Search(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, res.StatusCode)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d, err := strconv.Atoi(string(b))
	if err != nil {
		t.Fatalf("%v", err)
	}

	bs, err := ioutil.ReadFile("template/index.gohtml")
	if err != nil {
		t.Fatalf("%v", err)
	}
	dd, err := strconv.Atoi(string(bs))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if dd != d {
		t.Fatalf("expected %v but got %v", dd, d)
	}

}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000", nil)
	if err != nil {
		t.Fatalf("could not create the request: %v", err)
	}
	rec := httptest.NewRecorder()
	var s Server
	s.Index(rec, req)
	if status := rec.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}
