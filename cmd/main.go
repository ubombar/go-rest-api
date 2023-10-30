package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func greet(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Invoked header")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	logrus.Infoln("Setting up handlers")

	http.HandleFunc("/", greet)

	logrus.Infoln("Starting server")
	http.ListenAndServe(":80", nil)
	logrus.Infoln("Server shutting down")
}
