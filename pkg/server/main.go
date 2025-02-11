package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")

	log.Info("Recieved request on / -> responded with OK")
}

var log = logrus.New()

func Listen(address string) {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	log.Info("Server started at " + address)
	http.ListenAndServe(address, r)
}
