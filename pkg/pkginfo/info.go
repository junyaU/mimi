package pkginfo

import (
	"errors"
	"github.com/junyaU/mimi/pkg/utils"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"path/filepath"
	"strings"
)

type Package struct {
	Name    string
	Imports []string
}

type PackageOverview struct {
	Packages     []Package
	Dependencies []Package
	ModuleName   string
}

var (
	loadConfig = &packages.Config{
		Mode: packages.NeedImports | packages.NeedDeps | packages.NeedName,
	}
)

func New(root string) (*PackageOverview, error) {
	packagePaths, err := getPackages(root)
	if err != nil {
		return nil, err
	}

	moduleName, err := utils.GetModuleName()
	if err != nil {
		return nil, err
	}

	pkgOverview := PackageOverview{
		ModuleName: moduleName,
	}

	if err := loadPackages(&pkgOverview, packagePaths); err != nil {
		return nil, err
	}

	if err := loadDependencies(&pkgOverview); err != nil {
		return nil, err
	}

	return &pkgOverview, nil
}

func (p *PackageOverview) IsOwnPackage(pkgName string) bool {
	return strings.HasPrefix(pkgName, p.ModuleName)
}

func (p *PackageOverview) ExistsDependency(pkgName string) bool {
	for _, pkg := range p.Dependencies {
		if pkg.Name == pkgName {
			return true
		}
	}

	return false
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

func loadPackages(pkgOverview *PackageOverview, pkgPaths []string) error {
	if len(pkgPaths) == 0 {
		return errors.New("no packages")
	}

	pkgs, err := packages.Load(loadConfig, pkgPaths...)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		var imports []string
		for importPkg := range pkg.Imports {
			if pkgOverview.IsOwnPackage(importPkg) {
				imports = append(imports, importPkg)
			}
		}

		pkgOverview.Packages = append(pkgOverview.Packages, Package{
			Name:    pkg.ID,
			Imports: imports,
		})
	}

	return nil
}

func loadDependencies(pkgOverview *PackageOverview) error {
	for _, pkg := range pkgOverview.Packages {
		if err := insertDependencies(pkgOverview, pkg.Imports); err != nil {
			return err
		}
	}

	return nil
}

func insertDependencies(pkgOverview *PackageOverview, imports []string) error {
	loadPkgs, err := packages.Load(loadConfig, imports...)
	if err != nil {
		return err
	}

	for _, loadPkg := range loadPkgs {
		if pkgOverview.ExistsDependency(loadPkg.ID) {
			continue
		}

		var importPkgs []string
		importPkgs = make([]string, 0, len(loadPkg.Imports))
		for importPkg := range loadPkg.Imports {
			if pkgOverview.IsOwnPackage(importPkg) {
				importPkgs = append(importPkgs, importPkg)
			}
		}

		pkgOverview.Dependencies = append(pkgOverview.Dependencies, Package{
			Name:    loadPkg.ID,
			Imports: importPkgs,
		})

		if len(importPkgs) == 0 {
			continue
		}

		if err := insertDependencies(pkgOverview, importPkgs); err != nil {
			return err
		}
	}

	return nil
}
