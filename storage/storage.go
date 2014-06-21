package storage

import (
	"./lru"
	goconf "github.com/akrennmair/goconf"
	"log"
	"os"
)

type driver interface {
	Get(string) (string, bool)
	Set(string, string)
}

func driverFactory(config *goconf.ConfigFile) driver {
	drivername, err := config.GetString("storage", "driver")
	if err != nil {
		log.Printf("driver not setup")
		os.Exit(1)
	}
	switch drivername {
	case "memcached":
		return NewMemcached(config)
	default:
		log.Printf("driver %s not implemented", drivername)
		os.Exit(1)
	}
	return nil

}

type Storage struct {
	lru    *lru.LRU
	config *goconf.ConfigFile
	driver driver
}

func New(config *goconf.ConfigFile) *Storage {
	return &Storage{
		config: config,
		lru:    lru.New(10000),
		driver: driverFactory(config),
	}
}

func (c *Storage) Get(key string) (conexionstring string, ok bool) {
	conexionstring, r := c.lru.Get(key)
	if r {
		return conexionstring, true
	} else {
		conexionstring, r = c.driver.Get(key)
		if r {
			c.lru.Insert(key, conexionstring)
		}
		return conexionstring, r
	}
}
