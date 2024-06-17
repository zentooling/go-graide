package main

import (
	"net/http"

	"github.com/zentooling/graide/internal/config"
	"github.com/zentooling/graide/internal/logger"
	"github.com/zentooling/graide/internal/server"
	"github.com/zentooling/graide/internal/template"
)

type Contact struct {
	ID     int
	First  string
	Last   string
	Phone  string
	Email  string
	Errors map[string]string
}

func NewContact() Contact {
	return Contact{
		Errors: make(map[string]string),
	}
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
			Search: r.URL.Query().Get("q"),
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
			Search: r.URL.Query().Get("q"),
			Contacts: []Contact{
				{ID: 1, First: "Bill", Last: "Hambone", Phone: "303-555-1212", Email: "joey1@hambone.com"},
			},
		}
		err := view.ExecuteTemplate(w, template.CONTACT, data)
		if err != nil {
			log.Println("unable to execute template "+template.CONTACT, err)
		}
	})
	mux.HandleFunc("GET /contact/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		log.Printf("GET /contact/%s\n", idString)
		data := Contact{ID: 1, First: "Bill", Last: "Hambone", Phone: "303-555-1212", Email: "joey1@hambone.com"}
		err := view.ExecuteTemplate(w, template.EDIT, data)
		if err != nil {
			log.Println("unable to execute template "+template.EDIT, err)
		}
	})
	mux.HandleFunc("DELETE /contact/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("in DELETE /contact/%s ", r.PathValue("id"))
		log.Println("redirecting ...")
		http.Redirect(w, r, "/contact", http.StatusSeeOther) // 303 response to DELETE initiates GET request
	})
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contact", http.StatusFound)
	})
	mux.HandleFunc("GET /contact/new", func(w http.ResponseWriter, r *http.Request) {
		err := view.ExecuteTemplate(w, template.NEW, NewContact())
		if err != nil {
			log.Println("unable to execute template "+template.NEW, err)
		}
	})
	mux.HandleFunc("POST /contact/new", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for key, value := range r.Form {
			log.Printf("key %s value %s\n", key, value)
		}

		http.Redirect(w, r, "/contact", http.StatusFound)
	})
	err := mux.Listen()
	log.Fatal(err)

}
