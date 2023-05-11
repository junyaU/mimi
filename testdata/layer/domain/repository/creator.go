package repository

import "github.com/junyaU/mimi/testdata/layer/domain/model/creator"

type Creator interface {
	Get(id string) (*creator.Creator, error)
}
