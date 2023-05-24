package analysis

import (
	"github.com/junyaU/mimi/pkg/pkginfo"
	"testing"
)

const creatorPackage = "github.com/junyaU/mimi/testdata/layer/domain/model/creator"

func TestNewProjectPackages(t *testing.T) {
	projectPkgs, err := BuildTestProjectPackages(t)
	if err != nil {
		t.Errorf("NewProjectPackages() should not return error, but got %v", err)
	}

	if len(projectPkgs.Packages) != 1 {
		t.Errorf("NewProjectPackages() should return %v, but got %v", 1, len(projectPkgs.Packages))
	}

	if len(projectPkgs.Packages[0].Files) != 2 {
		t.Errorf("NewProjectPackages() should return %v, but got %v", 2, len(projectPkgs.Packages[0].Files))
	}
}

func TestProjectPackages_GetPackage(t *testing.T) {
	projectPkgs, err := BuildTestProjectPackages(t)
	if err != nil {
		t.Errorf("NewProjectPackages() should not return error, but got %v", err)
	}

	pkg, err := projectPkgs.GetPackage(creatorPackage)
	if err != nil {
		t.Errorf("GetPackage() should not return error, but got %v", err)
	}

	if pkg.Name != creatorPackage {
		t.Errorf("GetPackage() should return %v, but got %v", creatorPackage, pkg.Name)
	}
}

func TestPackageDetails_GetLines(t *testing.T) {
	projectPkgs, err := BuildTestProjectPackages(t)
	if err != nil {
		t.Errorf("NewProjectPackages() should not return error, but got %v", err)
	}

	pkg, err := projectPkgs.GetPackage(creatorPackage)
	if err != nil {
		t.Errorf("GetPackage() should not return error, but got %v", err)
	}

	if pkg.GetLines() != 33 {
		t.Errorf("GetLines() should return %v, but got %v", 33, pkg.GetLines())
	}
}

func BuildTestProjectPackages(t *testing.T) (*ProjectPackages, error) {
	t.Helper()

	info, err := pkginfo.New("./../../testdata/layer/domain/model/creator")
	if err != nil {
		return nil, err
	}

	projectPkgs, err := NewProjectPackages(info.Packages)
	if err != nil {
		return nil, err
	}

	return &projectPkgs, nil
}
