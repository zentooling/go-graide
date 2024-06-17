package main

import (
	"net/http"

	"github.com/zentooling/graide/internal/config"
	"github.com/zentooling/graide/internal/logger"
	"github.com/zentooling/graide/internal/server"
	"github.com/zentooling/graide/internal/template"
)

type Contact struct {
	ID    int
	First string
	Last  string
	Phone string
	Email string
}

type IndexPageData struct {
	Search   string
	Contacts []Contact
}

var log = logger.New("main")

func main() {

	log.Println("server startup")

	cfg := config.New("config.yml")

	log.Printf("+%v", *cfg)

	log.Println("template initialization")
	view := template.New()
	mux := server.New(cfg.Server.Host + ":" + cfg.Server.Port)
	mux.HandleFunc("GET /index", func(w http.ResponseWriter, r *http.Request) {
		data := IndexPageData{
			Search: "",
			Contacts: []Contact{
				{ID: 0, First: "Joey", Last: "Hambone", Phone: "303-555-1212", Email: "joey@hambone.com"},
			},
		}
		err := view.ExecuteTemplate(w, template.INDEX, data)
		if err != nil {
			log.Println("unable to execute template "+template.INDEX, err)
		}
	})
	mux.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		data := IndexPageData{
			Search: "",
			Contacts: []Contact{
				{ID: 1, First: "Bill", Last: "Hambone", Phone: "303-555-1212", Email: "joey1@hambone.com"},
			},
		}
		err := view.ExecuteTemplate(w, template.CONTACT, data)
		if err != nil {
			log.Println("unable to execute template "+template.CONTACT, err)
		}
	})
	err := mux.Listen()
	log.Fatal(err)

}
