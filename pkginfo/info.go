package pkginfo

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

type Info struct {
}

func New(root string) *Info {
	files, err := getFilePaths(root)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Info{}
}

func getFilePaths(root string) (files []string, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.Contains(d.Name(), ".go") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return
}
