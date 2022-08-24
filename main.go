package godiagnostics

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

func diagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("diagnosticsHandler")
}

func StartDiagnostics() {
	r := runtime.NumGoroutine()
	log.Println(r)

	router := mux.NewRouter()

	router.HandleFunc("/godiagnostics", diagnosticsHandler)

	http.ListenAndServe(":9000", router)
}
