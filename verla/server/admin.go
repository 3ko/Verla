package main

import (
	// "net"
	"net/http"
	// "os"
	// "sync/atomic"
)

type Admin struct {
	Storage    *Storage
	BindString string
}

func (a *Admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
