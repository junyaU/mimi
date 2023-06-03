package analysis

import (
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/junyaU/mimi/pkg/utils"
	"testing"
)

const (
	flowPackage = "github.com/junyaU/mimi/testdata/layer/domain/model/flow"
	testPath    = "./../../testdata/layer/domain/model/flow"
)

func TestNewGraph(t *testing.T) {
	graph := BuildDepGraph(t, testPath)

	if graph.nodes[0].Package != flowPackage {
		t.Errorf("NewGraph() should return %v, but got %v", flowPackage, graph.nodes[0].Package)
	}

	directDeps := []string{"github.com/junyaU/mimi/testdata/layer/domain", "github.com/junyaU/mimi/testdata/layer/domain/model/recipe"}
	for _, dep := range graph.nodes[0].Direct {
		if !utils.Contains(directDeps, dep) {
			t.Errorf("NewGraph() should return %v, but got %v", directDeps, graph.nodes[0].Direct)
		}
	}

	indirectDeps := []string{"github.com/junyaU/mimi/testdata/layer/domain/model/creator"}
	for _, dep := range graph.nodes[0].Indirect {
		if !utils.Contains(indirectDeps, dep) {
			t.Errorf("NewGraph() should return %v, but got %v", indirectDeps, graph.nodes[0].Indirect)
		}
	}
}

func TestPrintRows(t *testing.T) {
	tests := []struct {
		sortType      SortType
		expectLastRow []string
	}{
		{NoSort, []string{"github.com/junyaU/mimi/testdata/layer/infra", "5", "8", "0", "6", "0", "0.750000"}},
		{SortByWeight, []string{"github.com/junyaU/mimi/testdata/layer/adapter", "0", "0", "2", "0", "0", "0.050000"}},
	}

	for _, test := range tests {
		graph := BuildDepGraph(t, "./../../testdata")
		graph.AnalyzeIndirectDeps()
		graph.AnalyzeDependents()
		graph.AnalyzeWeights()

		result := graph.PrintRows(test.sortType)

		if result[len(result)-1][6] != test.expectLastRow[6] {
			t.Errorf("PrintRows() should return %v, but got %v", test.expectLastRow[6], result[len(result)-1][6])
		}
	}
}

func TestAnalyzeDirectDeps(t *testing.T) {
	info, err := pkginfo.New(testPath)
	if err != nil {
		t.Errorf("NewInfo() should not return error, but got %v", err)
	}

	graph := BuildDepGraph(t, testPath)

	if graph.nodes[0].Package != flowPackage {
		t.Errorf("NewGraph() should return %v, but got %v", flowPackage, graph.nodes[0].Package)
	}

	for _, dep := range graph.nodes[0].Direct {
		if !utils.Contains(info.Packages[0].Imports, dep) {
			t.Errorf("NewGraph() should return %v, but got %v", info.Packages[0].Imports, graph.nodes[0].Direct)
		}
	}
}

func TestAnalyzeIndirectDeps(t *testing.T) {
	graph := BuildDepGraph(t, testPath)
	graph.AnalyzeIndirectDeps()

	if graph.nodes[0].Package != flowPackage {
		t.Errorf("NewGraph() should return %v, but got %v", flowPackage, graph.nodes[0].Package)
	}

	if graph.nodes[0].Indirect[0] != "github.com/junyaU/mimi/testdata/layer/domain/model/creator" {
		t.Errorf("NewGraph() should return %v, but got %v", "github.com/junyaU/mimi/testdata/layer/domain/model/creator", graph.nodes[0].Indirect[0])
	}
}

func TestAnalyzeDependents(t *testing.T) {
	graph := BuildDepGraph(t, "./../../testdata/layer/domain/model/")
	graph.AnalyzeDependents()

	recipePkg := "github.com/junyaU/mimi/testdata/layer/domain/model/recipe"
	if graph.nodes[0].Dependents[0] != recipePkg {
		t.Errorf("NewGraph() should return %v, but got %v", recipePkg, graph.nodes[0].Dependents[0])
	}

	if graph.nodes[1].Dependents[0] != flowPackage {
		t.Errorf("NewGraph() should return %v, but got %v", flowPackage, graph.nodes[1].Dependents[0])
	}
}

func BuildDepGraph(t *testing.T, path string) *DepGraph {
	t.Helper()

	info, err := pkginfo.New(path)
	if err != nil {
		t.Fatalf("Failed to create pkginfo: %v", err)
	}

	graph, err := NewDepGraph(info)
	if err != nil {
		t.Fatalf("Failed to create depgraph: %v", err)
	}

	return graph
}
