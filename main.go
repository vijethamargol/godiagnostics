package godiagnostics

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

type Diagnostics struct {
	NumGoroutine int `json:"numGoroutine"`
}

func diagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("diagnosticsHandler")

	diagnostics := Diagnostics{
		NumGoroutine: runtime.NumGoroutine(),
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diagnostics)

}

func StartDiagnostics() {

	router := mux.NewRouter()

	router.HandleFunc("/godiagnostics", diagnosticsHandler)

	http.ListenAndServe(":9000", router)
}
