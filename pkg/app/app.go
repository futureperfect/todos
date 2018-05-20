package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	r *httprouter.Router
}

func (s *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func NewApp() *App {
	s := &App{
		r: httprouter.New(),
	}

	router := s.r

	router.GET("/", s.Index)

	return s
}

func (s *App) Run() {
	log.Printf("Serving on HTTP on %v", ":8080")

	log.Fatal(http.ListenAndServe(":8080", s.r))
}
