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

func (r *Server) Run() error {
	fmt.Println("connected to localhost: 8000")
	r.httphandler()
	if err := r.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *Server) httphandler() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Note Found 404\n"))
	})
	router.HandleFunc("/", r.Index).Methods("GET")
	router.HandleFunc("/search", r.Search).Methods("GET")
	router.HandleFunc("/search/{directory}/{keyword}", r.SearchProcess).Methods("GET")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	r.srv.Handler = router
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	s.tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.Redirect(w, r, "/search", 303)
}

func (s *Server) SearchProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// color.Red("======Request from http======")

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
