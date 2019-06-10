package main

import (
	"time"

	"github.com/syifan/goseth/server"
)

type counter struct {
	v uint64
}

func (c *counter) increase() {
	c.v++
}

func main() {
	c := counter{}
	s := server.NewServer(&c)
	go func(c *counter) {
		for {
			c.v++
			time.Sleep(100 * time.Millisecond)
		}
	}(&c)
	s.Run()
}
