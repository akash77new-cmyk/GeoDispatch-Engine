package models

// NodeID uniquely identifies a vertex in the road network graph.
type NodeID string

// Node represents an intersection or waypoint in the road network.
type Node struct {
	ID       NodeID   `json:"id"`
	Location Location `json:"location"`
}

// Edge represents a directed road segment connecting two nodes. Weight is
// the traversal cost (currently travel time in seconds; distance in
// meters is also tracked so routing and benchmarking can report both).
type Edge struct {
	To       NodeID  `json:"to"`
	Distance float64 `json:"distance_meters"`
	Weight   float64 `json:"weight_seconds"`
}

// Graph is an adjacency-list representation of the road network. It is
// the shared data structure that both Dijkstra and A* operate over.
//
// A plain map-of-slices is used instead of a matrix because road networks
// are sparse (each intersection connects to a handful of neighbors, not
// to every other node), so an adjacency list keeps memory and iteration
// cost proportional to the number of edges rather than nodes^2.
type Graph struct {
	Nodes map[NodeID]Node
	Adj   map[NodeID][]Edge
}

// NewGraph returns an empty, ready-to-use Graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[NodeID]Node),
		Adj:   make(map[NodeID][]Edge),
	}
}

// AddNode registers a node in the graph if it does not already exist.
func (g *Graph) AddNode(n Node) {
	if _, exists := g.Nodes[n.ID]; !exists {
		g.Nodes[n.ID] = n
	}
}

// AddEdge inserts a directed edge from -> to. Callers wanting an
// undirected road segment should call AddEdge twice (once per direction).
func (g *Graph) AddEdge(from NodeID, e Edge) {
	g.Adj[from] = append(g.Adj[from], e)
}
