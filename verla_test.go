package main

import (
	"./storage"
	goconf "github.com/akrennmair/goconf"
	"log"
	"testing"
	"time"
)

func TestMemcacheStorage(t *testing.T) {
	config := goconf.NewConfigFile()
	config.AddOption("storage", "driver", "memcached")
	config.AddOption("storage", "host", "127.0.0.1:11211")

	storage := storage.New(config)

	start := time.Now()
	storage.Set("a", "127.0.0.1:8080")
	elapsed := time.Since(start)
	log.Printf("Set storage took %s", elapsed)

	start = time.Now()
	storage.Get("a")
	elapsed = time.Since(start)
	log.Printf("Get storage took %s", elapsed)

	start = time.Now()
	storage.Get("a")
	elapsed = time.Since(start)
	log.Printf("Get from cache storage took %s", elapsed)
}

func BenchmarkStorageGetMencache(b *testing.B) {

	config := goconf.NewConfigFile()
	config.AddOption("storage", "driver", "memcached")
	config.AddOption("storage", "host", "127.0.0.1:11211")

	storage := storage.New(config)
	storage.Set("a", "127.0.0.1:8080")

	for i := 0; i < b.N; i++ {
		storage.Get("a")
	}
}
