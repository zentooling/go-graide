package server

import (
	"net/http"
	"sync"

	"github.com/zentooling/graide/internal/logger"
)

type ServerStruct struct {
	address string
	server  *http.ServeMux
}

var log = logger.New("server")

var mutex = sync.Mutex{}

var Server *ServerStruct = nil

func New(listenAddr string) *ServerStruct {
	mutex.Lock()
	defer mutex.Unlock()
	if Server == nil {
		mux := http.NewServeMux()
		fileServer := http.FileServer(http.Dir("../../web/static/"))
		mux.Handle("GET /web/static/", http.StripPrefix("/web/static", fileServer))
		Server = &ServerStruct{
			server:  mux,
			address: listenAddr,
		}
	}

	return Server

}

func Instance() *ServerStruct {
	if Server == nil {
		log.Fatalln("Server global not initialized. New() must be called before Instance()")
	}

	return Server
}

// blocking
func (s ServerStruct) Listen() error {
	return http.ListenAndServe(s.address, s.server)
}

func (s ServerStruct) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.server.HandleFunc(pattern, handler)
}
