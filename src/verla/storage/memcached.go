package storage

import (
	goconf "github.com/akrennmair/goconf"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
)

type Memcached struct {
	mc *memcache.Client
}

//
func NewMemcached(config *goconf.ConfigFile) *Memcached {
	host, err := config.GetString("storage", "host")
	if err != nil {
		log.Printf("host not set")
		os.Exit(1)
	}
	mc := memcache.New(host)
	//test if conection is up
	return &Memcached{
		mc: mc,
	}
}

func (c *Memcached) Get(key string) (conexionstring string, ok bool) {
	it, err := c.mc.Get(key)
	if err != nil {
		return "", false
	}
	return string(it.Value), true
}

func (c *Memcached) Set(key string, conexionstring string) error {
	return c.mc.Set(&memcache.Item{Key: key, Value: []byte(conexionstring)})
}

func (c *Memcached) Delete(key string) error {
	return c.mc.Delete(key)
}
