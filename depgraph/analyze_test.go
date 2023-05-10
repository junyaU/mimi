package depgraph

import (
	"github.com/junyaU/mimi/pkginfo"
	"log"
	"testing"
)

func TestNewGraph(t *testing.T) {
	info, err := pkginfo.New("./../testdata/layer/domain/model/flow")
	if err != nil {
		t.Errorf("NewInfo() should not return error, but got %v", err)
	}

	graph := New(info)
	log.Println(graph)
}
