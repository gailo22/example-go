package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {

	c := cache.New(24*time.Hour, 25*time.Hour)

	c.Set("foo", "bar", cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

}
