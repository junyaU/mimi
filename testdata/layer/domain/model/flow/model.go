package flow

import (
	"errors"
	"strings"
)

const _modelPath = "recipes/model/"

type model struct {
	dir       string
	name      string
	extension string
}

func newModel(name string) (m model, er error) {
	if err := m.register(name); err != nil {
		er = err
		return
	}

	if err := m.checkValidName(); err != nil {
		er = err
	}

	return
}

func (m model) showModelPath() string {
	return m.dir + m.name + m.extension
}

func (m *model) register(name string) error {
	index, err := m.checkExtension(name)
	if err != nil {
		return err
	}

	m.dir = _modelPath
	m.name = name[:index]
	m.extension = name[index:]

	return nil
}

func (m model) checkValidName() (err error) {
	if len(m.name) <= 0 {
		err = errors.New("有効なパス名ではありません")
	}
	return
}

func (m model) checkExtension(path string) (index int, err error) {
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
