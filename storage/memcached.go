package storage

import (
	goconf "github.com/akrennmair/goconf"
	// "github.com/bradfitz/gomemcache/memcache"
)

type Memcached struct {
}

//
func NewMemcached(config *goconf.ConfigFile) *Memcached {
	return &Memcached{}
}

func (c *Memcached) Get(key string) (conexionstring string, ok bool) {
	return "",false
}

func (c *Memcached) Set(key string, conexionstring string) {

}
//
// func main() {
// 	mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
// 	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
//
// 	it, err := mc.Get("foo")
// }
