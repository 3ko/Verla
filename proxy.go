package main

import (
	"log"
	// "net"
	"net/http"
	// "os"
)

type Backend struct {
	Name          string
	ConnectString string
}

type Frontend struct {
	Name         string
	BindString   string
	HTTPS        bool
	AddForwarded bool
	KeyFile      string
	CertFile     string
}

type RequestHandler struct {
	Transport    *http.Transport
	Frontend     *Frontend
	HostBackends map[string]chan *Backend
	Backends     chan *Backend
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("incoming request: %#v", *r)
}
