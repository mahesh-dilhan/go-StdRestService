package greethandler

import (
	"log"
	"net/http"
)

type Greet struct {
	l *log.Logger
}

func NewGreet(l *log.Logger) *Greet {
	return &Greet{l}
}

func (g *Greet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello world")
}
