package main

import (
	"./storage"
	"flag"
	"fmt"
	goconf "github.com/akrennmair/goconf"
	// "net/http"
	"log"
	"os"
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

	// cache := lru.New(300)
	// cache.Insert("test.fr", "127.0.0.1:8000")
	// val, ok := cache.Get("test.fr")
	// fmt.Println(ok, val)

	var cfgfile *string = flag.String("config", "", "configuration file")

	flag.Parse()

	if *cfgfile == "" {
		usage()
	}

	config, err := goconf.ReadConfigFile(*cfgfile)

	if err != nil {
		log.Printf("opening %s failed: %v", *cfgfile, err)
		os.Exit(1)
	}

	storage := storage.New(config)
	// storage.Set("a", "127.0.0.1:8080")
	storage.Get("a")
	// log.Printf(config)
	// hosts_chans := make(map[string]chan *Backend)
	// backends_chan := make(chan *Backend)
	// mux := http.NewServeMux()
	// var request_handler http.Handler = &RequestHandler{Transport: &http.Transport{DisableKeepAlives: false, DisableCompression: false}, HostBackends: hosts_chans, Backends: backends_chan}
	// mux.Handle("/", request_handler)
	// srv := &http.Server{Handler: mux, Addr: ":8080"}
}
