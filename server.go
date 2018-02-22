package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Tomoka64/final/model"
	"github.com/gorilla/mux"
)

//Server consists of http.Server and template.
type Server struct {
	srv *http.Server
	tpl template.Template
}

func newServer(items ...string) (another, error) {

	hostname := "127.0.0.1:8000"
	srv := &http.Server{
		Addr:         hostname,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &Server{
		srv: srv,
		tpl: *template.Must(template.ParseGlob("template/*")),
	}, nil
}

//Run runs the server, using httphandler.
func (s *Server) Run() error {
	fmt.Println("connected to localhost: 8000")
	s.httphandler()
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}

//httphandler implements gorilla mux.
func (s *Server) httphandler() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Found 404\n"))
	})
	router.HandleFunc("/", s.Index).Methods("GET")
	router.HandleFunc("/search", s.Search).Methods("GET")
	router.HandleFunc("/search/{directory}/{keyword}", s.SearchProcess).Methods("GET")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	s.srv.Handler = router
}

//Search will serve Template index.gohtml
func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	s.tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

//Index will redirect to /search when the http request is "/"
func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.Redirect(w, r, "/search", 303)
}

//SearchProcess gets parameters from URL (/search/{directory}/{keyword}) and
//searches the package and word accordingly and writes it in json format.
func (s *Server) SearchProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	params := mux.Vars(r)
	path := params["directory"]
	p, err := importPkg(path, dir)
	if err != nil {
		log.Fatalln(err)
	}

	pattern := params["keyword"]
	datas := []model.Result{}
	for _, f := range p.GoFiles {
		fname := filepath.Join(p.Dir, f)
		datas, err = extractWord(fname, pattern, datas)
		if err != nil {
			break
		}
	}

	json.NewEncoder(w).Encode(datas)
}
