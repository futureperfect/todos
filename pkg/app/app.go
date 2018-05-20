package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	r *httprouter.Router
}

type Todo struct {
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Due         time.Time `json:"due"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func (s *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (s *App) ListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todos: []Todo{Todo{Description: "Write a server"},
			Todo{Description: "Deploy it!"}},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func NewApp() *App {
	s := &App{
		r: httprouter.New(),
	}

	router := s.r

	router.GET("/todos", s.ListTodos)
	router.GET("/", s.Index)

	return s
}

func (s *App) Run() {
	log.Printf("Serving on HTTP on %v", ":8080")

	log.Fatal(http.ListenAndServe(":8080", s.r))
}
