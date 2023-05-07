package main

import (
	"github.com/junyaU/mimi/depgraph"
	"github.com/junyaU/mimi/pkginfo"
	"log"
)

func main() {
	info, err := pkginfo.New("./testdata")
	if err != nil {
		panic(err)
	}

	aaa := depgraph.New(info)

	log.Println(aaa)
}
