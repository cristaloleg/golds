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

// LabeledEdge ...
type LabeledEdge interface {
	Edge
	Label() string
}

// WeightedEdge ...
type WeightedEdge interface {
	Edge
	Weight() float64
}

// DirectedWeightedEdge ...
type DirectedWeightedEdge interface {
	Edge
	Source() Vertex
	Target() Vertex
	Weight() float64
}

type edge struct {
	u, v Vertex
}

type directedEdge struct {
	Edge
}

type labeledEdge struct {
	Edge
	label string
}

type weightedEdge struct {
	Edge
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
	e := directedEdge{
		Edge: edge{u: u, v: v},
	}
	return e
}

func (e directedEdge) Vertices() (u, v Vertex) {
	return e.Vertices()
}

func (e directedEdge) Source() (u Vertex) {
	u, _ = e.Vertices()
	return u
}

func (e directedEdge) Target() (v Vertex) {
	_, v = e.Vertices()
	return v
}

// NewLabeledEdge ...
func NewLabeledEdge(u, v Vertex, s string) LabeledEdge {
	e := labeledEdge{
		Edge:  edge{u, v},
		label: s,
	}
	return e
}

func (e labeledEdge) Vertices() (u, v Vertex) {
	return e.Vertices()
}

func (e labeledEdge) Label() string {
	return e.label
}

// NewWeightedEdge ...
func NewWeightedEdge(u, v Vertex, w float64) WeightedEdge {
	e := weightedEdge{
		Edge:   edge{u, v},
		weight: w,
	}
	return e
}

func (e weightedEdge) Vertices() (u, v Vertex) {
	return e.Vertices()
}

func (e weightedEdge) Weight() float64 {
	return e.weight
}

// Graph ...
type Graph interface {
	Has(Vertex) bool
	HasE(Edge) bool
	Connected(a, b Vertex) bool
	EdgesOf(Vertex) []Edge

	Edges() []Edge
	Vertices() []Vertex

	AddEdge(Edge)
	DelEdge(Edge)
}

// DirectedGraph ...
type DirectedGraph interface {
	Graph
	InDegree(Vertex) int
	OutDegree(Vertex) int
}

// WeightedGraph ...
type WeightedGraph interface {
	Graph
	Sum(Vertex) float64
}

// Type ...
type Type int

const (
	// Simple ...
	Simple Type = 1 << iota
	// Loop ...
	Loop
	// Parallel ...
	Parallel

	// Directed ...
	Directed
	// Labeled ...
	Labeled
	// Weighted ...
	Weighted
)

// Build ...
func Build(t Type) Graph {
	return nil
}

var _ Graph = Build(Simple | Loop | Parallel | Directed | Labeled | Weighted)
