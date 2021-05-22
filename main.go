package main

import (
	"github.com/mahesh-dilhan/go-StdRestService/greethandler"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "covidservice-api", log.LstdFlags)
	gh := greethandler.NewGreet(l)

	sm := http.NewServeMux()
	sm.Handle("/", gh)
	http.ListenAndServe(":9090", sm)
}
