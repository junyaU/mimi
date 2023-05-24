package analysis

import (
	"bufio"
	"fmt"
	"github.com/junyaU/mimi/pkg/pkginfo"
	"os"
)

type ProjectPackages struct {
	Packages []PackageDetails
}

type PackageDetails struct {
	Name  string
	Files []File
}

type File struct {
	Name  string
	Lines int
}

func NewProjectPackages(pkgs []pkginfo.Package) (ProjectPackages, error) {
	if len(pkgs) == 0 {
		return ProjectPackages{}, fmt.Errorf("no packages found")
	}

	var projectPackages ProjectPackages

	for _, pkg := range pkgs {
		detail := PackageDetails{
			Name: pkg.Name,
		}

		for _, f := range pkg.Files {
			file, err := os.Open(f)
			if err != nil {
				return ProjectPackages{}, fmt.Errorf("failed to open file %s: %w", f, err)
			}

			scanner := bufio.NewScanner(file)
			lines := 0
			for scanner.Scan() {
				lines++
			}

			if err := scanner.Err(); err != nil {
				return ProjectPackages{}, fmt.Errorf("failed to scan file %s: %w", f, err)
			}

			detail.Files = append(detail.Files, File{
				Name:  file.Name(),
				Lines: lines,
			})

			file.Close()
		}

		projectPackages.Packages = append(projectPackages.Packages, detail)
	}

	return projectPackages, nil
}

func (p *ProjectPackages) GetPackage(pkgName string) (PackageDetails, error) {
	for _, pkg := range p.Packages {
		if pkg.Name == pkgName {
			return pkg, nil
		}
	}

	return PackageDetails{}, fmt.Errorf("package not found: %s", pkgName)
}

func (d *PackageDetails) GetLines() int {
	var lines int
	for _, f := range d.Files {
		lines += f.Lines
	}
	return lines
}
