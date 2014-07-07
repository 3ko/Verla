package main

import (
	"./proxy"
	"./storage"
	"flag"
	"fmt"
	goconf "github.com/akrennmair/goconf"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
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
	resquest_handler := proxy.New(storage)

	// genSelfCert("www.3ko.fr")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(*addr, resquest_handler))
		wg.Done()
	}()

	// go func() {
	// 	log.Fatal(http.ListenAndServeTLS(":82", resquest_handler))
	// 	wg.Done()
	// }()

	wg.Wait()
}