package domain

import (
	"encoding/json"
	"time"
)

type EventModel struct {
	Id        string `dynamo:",hash"`
	Version   uint64 `dynamo:",range"`
	CreatedAt time.Time
	EventType string
	EventData json.RawMessage
}

func NewEventModel(id string, version uint64, eventType string, eventData json.RawMessage) EventModel {
	return EventModel{
		Id:        id,
		Version:   version,
		EventType: eventType,
		EventData: eventData,
		CreatedAt: time.Now(),
	}
}

func (m EventModel) GetVersion() uint64 {
	return m.Version
}
