package creator

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

type Id string

func newId() Id {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return Id(id.String())
}
