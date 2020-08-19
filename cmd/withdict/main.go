package main

import (
	"bytes"
	"log"

	"github.com/ilius/stardict/pkg/tui"
)

func main() {
	// using this tool to convert data to golang source code
	// https://github.com/jteeuwen/go-bindata
	// the dict I'm using is contained in go source code with command:
	// go-bindata -o dict.go langdao-ec-gb.dict.dz langdao-ec-gb.idx langdao-ec-gb.ifo
	a, err := Asset("langdao-ec-gb.ifo")
	if err != nil {
		log.Fatal(err)
	}
	b, err := Asset("langdao-ec-gb.idx")
	if err != nil {
		log.Fatal(err)
	}
	c, err := Asset("langdao-ec-gb.dict.dz")

	r1 := bytes.NewReader(a)
	r2 := bytes.NewReader(b)
	r3 := bytes.NewReader(c)

	engine, err := tui.NewEngine(r1, r2, r3)
	if err != nil {
		log.Fatal(err)
	}
	engine.RunWithOutput()
}
