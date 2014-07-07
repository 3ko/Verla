package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	// "os"
	"../storage"
	"regexp"
	"strings"
	// "sync/atomic"
)

type Backend struct {
	Name          string
	ConnectString string
}

type Error struct {
	StatusCode int
	Message    string
}

type Frontend struct {
	Storage       *storage.Storage
	Name          string
	BindString    string
	HTTPS         bool
	lastRequestId int64
}

type Request struct {
	HttpRequest *http.Request
	Id          int64
	// Body netutils.MultiReader
}

var hasPort = regexp.MustCompile(`:\d+$`)

func (proxy *Frontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	host, ok := proxy.getBackend(r)

	if !ok {
		return
	}

	if isWebsocket(r) {
		p := httpproxy(host) //example URL
		p.ServeHTTP(w, r)
		return
	}

	if isHTTPS(r) {
		return
	}

	p := httpproxy(host)
	p.ServeHTTP(w, r)

}

func (proxy *Frontend) getBackend(req *http.Request) (string, bool) {
	return proxy.Storage.Get(req.Host)
}

func (proxy *Frontend) proxyRequest(w http.ResponseWriter, req *http.Request) *Error {
	addr, ok := proxy.Storage.Get(req.Host)
	if !ok {
		return &Error{StatusCode: http.StatusBadGateway}
	}
	log.Printf("%v", addr)
	return &Error{StatusCode: 500}
}

func New(storage *storage.Storage) *Frontend {
	frontend := Frontend{
		Storage: storage,
	}
	return &frontend
}

func httpproxy(target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := net.Dial("tcp", target)
		if err != nil {
			http.Error(w, "Error contacting backend server.", 500)
			log.Printf("Error dialing websocket backend %s: %v", target, err)
			return
		}
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "Not a hijacker?", 500)
			return
		}
		nc, _, err := hj.Hijack()
		if err != nil {
			log.Printf("Hijack error: %v", err)
			return
		}
		defer nc.Close()
		defer d.Close()

		err = r.Write(d)
		if err != nil {
			log.Printf("Error copying request to target: %v", err)
			return
		}

		errc := make(chan error, 2)
		cp := func(dst io.Writer, src io.Reader) {
			_, err := io.Copy(dst, src)
			errc <- err
		}
		go cp(d, nc)
		go cp(nc, d)
		<-errc
	})
}

func isHTTPS(req *http.Request) bool {
	return (req.Method == "CONNECT")
}

func isWebsocket(req *http.Request) bool {
	conn_hdr := ""
	conn_hdrs := req.Header["Connection"]
	if len(conn_hdrs) > 0 {
		conn_hdr = conn_hdrs[0]
	}

	upgrade_websocket := false
	if strings.ToLower(conn_hdr) == "upgrade" {
		upgrade_hdrs := req.Header["Upgrade"]
		if len(upgrade_hdrs) > 0 {
			upgrade_websocket = (strings.ToLower(upgrade_hdrs[0]) == "websocket")
		}
	}

	return upgrade_websocket
}
