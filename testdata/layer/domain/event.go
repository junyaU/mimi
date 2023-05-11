//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../_mock/$GOPACKAGE/$GOFILE
package domain

type Eventer interface {
	Identifier() string
	GetEventType() string
	CheckTheVersion() uint64
	SetVersion(version uint64)
}

type Event struct {
	id        string
	version   uint64
	eventType string
}

func NewEvent(id string, version uint64, eventType string) *Event {
	e := new(Event)
	e.id = id
	e.version = version
	e.eventType = eventType

	return e
}

func (e *Event) Identifier() string {
	return e.id
}

func (e *Event) GetEventType() string {
	return e.eventType
}

func (e *Event) CheckTheVersion() uint64 {
	return e.version
}

func (e *Event) SetVersion(version uint64) {
	e.version = version
}
