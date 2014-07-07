package admin

import (
	"net"
	"net/http"
	// "os"
	"../storage"
	// "sync/atomic"
)

type Admin struct {
	Storage    *storage.Storage
	BindString string
}

func (a *Admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
