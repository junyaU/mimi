package pkginfo

import (
	"github.com/junyaU/mimi/utils"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{"../testdata", false},
		{"../testdata/invalid", true},
	}

	for _, test := range tests {
		_, err := New(test.path)
		if err != nil && !test.wantErr {
			t.Error(err)
		}

		if err == nil && test.wantErr {
			t.Errorf("New(%v) should return error", test.path)
		}
	}
}

func TestGetPackages(t *testing.T) {
	tests := []struct {
		path    string
		expect  []string
		wantErr bool
	}{
		{"../testdata/layer/domain/model",
			[]string{
				"./../testdata/layer/domain/model/creator",
				"./../testdata/layer/domain/model/flow",
				"./../testdata/layer/domain/model/necessity",
				"./../testdata/layer/domain/model/recipe",
			},
			false},
		{"../testdata/invalid", nil, true},
	}

	for _, test := range tests {
		packagesPath, err := getPackages(test.path)
		if err != nil && !test.wantErr {
			t.Error(err)
		}

		if err == nil && test.wantErr {
			t.Errorf("getPackages(%v) should return error", test.path)
		}

		if !test.wantErr {
			for _, expect := range test.expect {
				if !utils.Contains(packagesPath, expect) {
					t.Errorf("getPackages(%v) should return %v", test.path, test.expect)
				}
			}
		}
	}
}

func TestLoadPackages(t *testing.T) {
	tests := []struct {
		pkgPaths  []string
		importPkg []string
		wantErr   bool
	}{
		{
			[]string{"../testdata/layer/domain/model/flow"},
			[]string{
				"github.com/junyaU/mimi/testdata/layer/domain",
				"github.com/junyaU/mimi/testdata/layer/domain/model/recipe",
			},
			false,
		},
		{
			[]string{},
			[]string{},
			true,
		},
	}

	for _, test := range tests {
		pkgInfo, err := loadPackages(test.pkgPaths)
		if err != nil && !test.wantErr {
			t.Error(err)
		}

		if err == nil && test.wantErr {
			t.Errorf("loadPackages(%v) should return error", test.pkgPaths)
		}

		if !test.wantErr {
			if pkgInfo[0].Name != "github.com/junyaU/mimi/testdata/layer/domain/model/flow" {
				t.Errorf("loadPackages(%v) should return %v", test.pkgPaths, test.importPkg)
			}

			for _, importPkg := range test.importPkg {
				log.Println(importPkg)
				log.Println(pkgInfo[0].Imports)
				if !utils.Contains(pkgInfo[0].Imports, importPkg) {
					t.Errorf("loadPackages(%v) should return %v", test.pkgPaths, test.importPkg)
				}
			}
		}
	}
}
