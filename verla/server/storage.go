package main

import (
	goconf "github.com/akrennmair/goconf"
	"log"
	"os"
)

type driver interface {
	Get(string) (string, bool)
	Set(string, string) error
	Delete(string) error
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
	lru    *LRU
	config *goconf.ConfigFile
	driver driver
}

func NewStorage(config *goconf.ConfigFile) *Storage {
	return &Storage{
		config: config,
		lru:    NewLRU(10000),
		driver: driverFactory(config),
	}
}

func (c *Storage) Get(key string) (conexionstring string, ok bool) {

	// log.Printf("get addr from %s ", key)

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

func (c *Storage) Set(key string, conexionstring string) {
	c.driver.Set(key, conexionstring)
}

func (c *Storage) Delete(key string) {
	c.lru.Delete(key)
	c.driver.Delete(key)
}
