package recipe

import (
	"github.com/junyaU/mimi/testdata/layer/domain/model/creator"
)

type Recipe struct {
	id          Id
	creatorID   creator.Id
	title       title
	thumbnail   thumbnail
	description description
	cookingTime cookingTime
	serving     serving
	status      status
}

func New(creatorId, title, thumbnail, description string, cookingTime, serving int) (*Recipe, error) {
	ti, err := newTitle(title)
	if err != nil {
		return nil, err
	}

	th, err := newThumbnail(thumbnail)
	if err != nil {
		return nil, err
	}

	de, err := newDescription(description)
	if err != nil {
		return nil, err
	}

	co, err := newCookingTime(cookingTime)
	if err != nil {
		return nil, err
	}

	se, err := newServing(serving)
	if err != nil {
		return nil, err
	}

	return &Recipe{
		id:          newId(),
		creatorID:   creator.Id(creatorId),
		title:       ti,
		thumbnail:   th,
		description: de,
		cookingTime: co,
		serving:     se,
		status:      PRIVATE,
	}, nil
}

func (r Recipe) Identifier() string {
	return string(r.id)
}

func (r Recipe) CreatorIdentifier() string {
	return string(r.creatorID)
}

func (r Recipe) Title() string {
	return r.title.String()
}

func (r Recipe) ThumbnailPath() string {
	return r.thumbnail.String()
}

func (r Recipe) Description() string {
	return r.description.String()
}

func (r Recipe) CookingTime() int {
	return int(r.cookingTime)
}

func (r Recipe) Serving() int {
	return int(r.serving)
}

func (r Recipe) Status() string {
	return r.status.String()
}
