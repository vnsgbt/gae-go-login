package main

import (
	"google.golang.org/appengine"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
)

func main()  {
	registerHandlers()
	appengine.Main()
}

func registerHandlers() {
	r := mux.NewRouter()
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

// http://blog.golang.org/error-handling-and-go
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}