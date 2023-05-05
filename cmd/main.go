package main

import (
	"github.com/junyaU/mimi/pkginfo"
)

func main() {
	//config := packages.Config{
	//	Mode:  packages.NeedDeps | packages.NeedImports,
	//	Tests: false,
	//}

	pkginfo.New("../Yur")

	//pkgs, err := packages.Load(&config, "../Yur/main.go")
	//if err != nil {
	//	log.Println(err)
	//	panic(err)
	//}

	//log.Println(pkgs[0].ExportFile)
}
