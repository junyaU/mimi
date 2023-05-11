package data_handler

import (
	"errors"
	"time"

	"github.com/junyaU/mimi/testdata/layer/adapter"
	"github.com/junyaU/mimi/testdata/layer/domain/model/creator"
)

type Creator struct {
	handler adapter.DataHandler
}

type creatorData struct {
	Id        string
	Name      string
	Email     string
	SelfIntro string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(h adapter.DataHandler) *Creator {
	return &Creator{
		handler: h,
	}
}

func (u *Creator) Get(id string) (*creator.Creator, error) {
	sql := "SELECT id, name, email, self_intro, icon, created_at, updated_at FROM users WHERE id = ?"
	var data creatorData
	if err := u.handler.Query(&data, sql, id); err != nil {
		return nil, err
	}

	if data.Id == "" {
		return nil, errors.New("no applicable data exists")
	}

	c, err := creator.New(data.Id, data.Name, data.Icon, data.SelfIntro)
	if err != nil {
		return nil, err
	}

	return c, nil
}
