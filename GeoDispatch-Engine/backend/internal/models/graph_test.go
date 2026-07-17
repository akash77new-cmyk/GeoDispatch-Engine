package models

import "testing"

// TestNewGraphIsEmpty ensures a freshly constructed graph has no nodes
// or edges, and that its maps are initialized (not nil) so callers can
// write to them immediately.
func TestNewGraphIsEmpty(t *testing.T) {
	g := NewGraph()

	if len(g.Nodes) != 0 {
		t.Errorf("expected 0 nodes, got %d", len(g.Nodes))
	}
	if len(g.Adj) != 0 {
		t.Errorf("expected 0 adjacency entries, got %d", len(g.Adj))
	}
}

// TestAddNodeIdempotent verifies that adding the same node ID twice does
// not overwrite the first insertion or create duplicate bookkeeping.
func TestAddNodeIdempotent(t *testing.T) {
	g := NewGraph()
	n := Node{ID: "n1", Location: Location{Lat: 1, Lng: 1}}

	g.AddNode(n)
	g.AddNode(Node{ID: "n1", Location: Location{Lat: 99, Lng: 99}})

	if g.Nodes["n1"].Location.Lat != 1 {
		t.Errorf("expected first insertion to be preserved, got lat=%f", g.Nodes["n1"].Location.Lat)
	}
}

// TestAddEdgeAppendsToAdjacencyList verifies edges accumulate under the
// correct source node.
func TestAddEdgeAppendsToAdjacencyList(t *testing.T) {
	g := NewGraph()
	g.AddEdge("a", Edge{To: "b", Distance: 100, Weight: 12})
	g.AddEdge("a", Edge{To: "c", Distance: 200, Weight: 24})

	edges := g.Adj["a"]
	if len(edges) != 2 {
		t.Fatalf("expected 2 edges, got %d", len(edges))
	}
	if edges[0].To != "b" || edges[1].To != "c" {
		t.Errorf("unexpected edge order: %+v", edges)
	}
}
