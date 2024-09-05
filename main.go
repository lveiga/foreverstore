package main

import (
	"log"

	p2p "github.com/desenv/foreverstore/p2p/tcp"
)

func main() {
	tr := p2p.NewTCPTransport(":3001")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
