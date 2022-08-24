package godiagnostics

import (
	"log"
	"runtime"
)

func StartDiagnostics() {
	r := runtime.NumGoroutine()
	log.Println(r)
}
