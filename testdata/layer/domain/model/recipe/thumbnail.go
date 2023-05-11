package recipe

import (
	"errors"
	"strings"
)

type thumbnail struct {
	dir       string
	name      string
	extension string
}

const _thumbnailDir = "recipes/thumbnail/"

func newThumbnail(path string) (t thumbnail, er error) {
	if err := t.register(path); err != nil {
		er = err
		return
	}

	if err := t.checkValidPath(); err != nil {
		er = err
	}

	return
}

func (t thumbnail) showThumbnailPath() string {
	return t.dir + t.name + t.extension
}

func (t *thumbnail) register(name string) error {
	index, err := t.checkExtension(name)
	if err != nil {
		return err
	}

	t.dir = _thumbnailDir
	t.name = name[:index]
	t.extension = name[index:]

	return nil
}

func (t thumbnail) checkValidPath() (err error) {
	if len(t.name) == 0 {
		err = errors.New("pathに値が設定されていません")
	}

	return
}

func (t thumbnail) checkExtension(path string) (index int, err error) {
	jpgIndex := strings.Index(path, ".jpg")
	jpegIndex := strings.Index(path, ".jpeg")
	pngIndex := strings.Index(path, ".png")

	switch {
	case jpgIndex != -1:
		index = jpgIndex
	case jpegIndex != -1:
		index = jpegIndex
	case pngIndex != -1:
		index = pngIndex
	default:
		err = errors.New("正しい拡張子ではありません")
	}

	return
}

func (t thumbnail) String() string {
	return t.dir + t.name + t.extension
}
