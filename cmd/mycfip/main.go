package main

import (
	"fmt"
	"log"

	"github.com/jpillora/ipflare"
	"github.com/jpillora/opts"
)

var version = "0.0.0-src"

func main() {
	c := struct{}{}
	opts.New(&c).Version(version).Parse()
	ip, err := ipflare.My()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip.String())
}
