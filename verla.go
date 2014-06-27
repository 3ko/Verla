package main

import (
	"./proxy"
	"./storage"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	goconf "github.com/akrennmair/goconf"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

type test struct {
	int
	string
}

func usage() {
	fmt.Fprintf(os.Stdout, "usage: %s -config=<configfile>\n", os.Args[0])
	os.Exit(1)
}

func main() {

	cfgfile := flag.String("config", "", "configuration file")
	addr := flag.String("addr", ":8080", "proxy listen address")

	flag.Parse()

	if *cfgfile == "" {
		usage()
	}

	config, err := goconf.ReadConfigFile(*cfgfile)
	threads, err := config.GetInt("default", "thread")
	if threads > 1 {
		runtime.GOMAXPROCS(threads)
	}
	if err != nil {
		log.Printf("opening %s failed: %v", *cfgfile, err)
		os.Exit(1)
	}

	storage := storage.New(config)
	storage.Delete("localhost:8080")
	storage.Set("localhost:8080", "game-synergy.fr:80")
	handler := proxy.New(storage)

	genSelfCert("www.3ko.fr")

	// http.ListenAndServeTLS(*addr2, "cert.pem", "key.pem", handler)
	http.ListenAndServe(*addr, handler)

}
