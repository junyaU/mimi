package analysis

import (
	"github.com/junyaU/mimi/pkg/pkginfo"
	"github.com/junyaU/mimi/pkg/utils"
	"strconv"
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
	graph := BuildDepGraph(t, testPath)
	graph.AnalyzeIndirectDeps()

	rows := graph.PrintRows()

	if len(rows[0]) != 5 {
		t.Errorf("PrintRows() should return %v, but got %v", 4, len(rows))
	}

	if rows[0][0] != flowPackage {
		t.Errorf("PrintRows() should return %v, but got %v", flowPackage, rows[0][0])
	}

	directDepsNum, err := strconv.Atoi(rows[0][1])
	if err != nil {
		t.Errorf("Num of direct dependencies should be integer, but got %v", rows[0][1])
	}

	if directDepsNum != 2 {
		t.Errorf("PrintRows() should return %v, but got %v", 2, directDepsNum)
	}

	indirectDepsNum, err := strconv.Atoi(rows[0][2])
	if err != nil {
		t.Errorf("Num of indirect dependencies should be integer, but got %v", rows[0][2])
	}

	if indirectDepsNum != 1 {
		t.Errorf("PrintRows() should return %v, but got %v", 1, indirectDepsNum)
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
