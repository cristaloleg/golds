package graph

// Vertex ...
type Vertex interface{}

// Edge ...
type Edge interface {
	Vertices() (u, v Vertex)
}

// DirectedEdge ...
type DirectedEdge interface {
	Edge
	Source() Vertex
	Target() Vertex
}

// WeightedEdge ...
type WeightedEdge interface {
	Edge
	Weight() float64
}

type edge struct {
	u, v Vertex
}

type directedEdge struct {
	u, v Vertex
}

type weightedEdge struct {
	u, v   Vertex
	weight float64
}

// NewEdge ...
func NewEdge(u, v Vertex) Edge {
	e := edge{u: u, v: v}
	return e
}

func (e edge) Vertices() (u, v Vertex) {
	return e.u, e.v
}

// NewDirectedEdge ...
func NewDirectedEdge(u, v Vertex) DirectedEdge {
	e := directedEdge{u: u, v: v}
	return e
}

func (e directedEdge) Vertices() (u, v Vertex) {
	return e.u, e.v
}

func (e directedEdge) Source() (u Vertex) {
	return e.u
}

func (e directedEdge) Target() (v Vertex) {
	return e.v
}

// NewWeightedEdge ...
func NewWeightedEdge(u, v Vertex, w float64) WeightedEdge {
	e := weightedEdge{u: u, v: v, weight: w}
	return e
}

func (e weightedEdge) Vertices() (u, v Vertex) {
	return e.u, e.v
}

func (e weightedEdge) Weight() float64 {
	return e.weight
}

// Graph ...
type Graph interface {
	Edges() int
	Vertices() int
}

// DirectedGraph ...
type DirectedGraph interface {
	AddDirectedEdge(DirectedEdge)
	HasDirectedEdge(DirectedEdge) bool
	InDegree(Vertex) int
	OutDegree(Vertex) int
}

// WeightedGraph ...
type WeightedGraph interface {
	AddWeightedEdge(WeightedEdge)
	HasWeightedEdge(WeightedEdge) bool
}
