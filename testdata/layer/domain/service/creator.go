package service

import (
	"github.com/junyaU/mimi/testdata/layer/domain/repository"
)

type Creator struct {
	repo repository.Creator
}

func NewCreatorService(r repository.Creator) *Creator {
	return &Creator{
		repo: r,
	}
}

func (c *Creator) Exists(id string) bool {
	creator, err := c.repo.Get(id)
	return creator != nil && err == nil
}
