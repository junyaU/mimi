package main

import (
	"github.com/junyaU/mimi/pkginfo"
	"log"
)

func main() {
	infos, err := pkginfo.New("./")
	if err != nil {
		panic(err)
	}

	log.Println("$$$$$$$")
	log.Println(infos)
}
