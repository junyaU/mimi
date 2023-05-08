package pkginfo

import (
	"github.com/junyaU/mimi/utils"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"path/filepath"
	"strings"
)

type Info struct {
	Name    string
	Imports []string
}

func New(root string) ([]Info, error) {
	packagePaths, err := getPackages(root)
	if err != nil {
		return nil, err
	}

	return loadPackages(packagePaths)
}

func getPackages(root string) (pkgs []string, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.Contains(d.Name(), ".go") {
			return nil
		}

		packageName := filepath.Dir(path)
		packagePattern := "./" + filepath.ToSlash(packageName)
		if pkgs == nil {
			pkgs = append(pkgs, packagePattern)
			return nil
		}

		if packagePattern != pkgs[len(pkgs)-1] {
			pkgs = append(pkgs, packagePattern)
		}

		return nil
	})

	return
}

func loadPackages(pkgPaths []string) ([]Info, error) {
	config := packages.Config{
		Mode:  packages.NeedDeps | packages.NeedImports | packages.NeedName,
		Tests: false,
	}

	pkgs, err := packages.Load(&config, pkgPaths...)
	if err != nil {
		return nil, err
	}

	moduleName, err := utils.GetModuleName()
	if err != nil {
		return nil, err
	}

	var infos []Info

	for _, pkg := range pkgs {
		var imports []string
		for k, _ := range pkg.Imports {
			if strings.Contains(k, moduleName) {
				imports = append(imports, k)
			}
		}

		infos = append(infos, Info{
			Name:    pkg.ID,
			Imports: imports,
		})
	}

	return infos, nil
}
