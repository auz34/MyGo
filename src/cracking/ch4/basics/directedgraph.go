package basics

type Vertex struct {
	Id int
	FromEdges map[int]*Edge
	ToEdges map[int]*Edge
}

type Edge struct {
	From int
	To int
}

type DirectedGraph struct {
	Vertices map[int]*Vertex
}

func (self *DirectedGraph) AddVertex(id int) {
	self.Vertices[id] = &Vertex{ id, make(map[int]*Edge), make(map[int]*Edge)}
}

func (self *DirectedGraph) AddEdge(from, to int) {
	edge := &Edge{ from, to }
	fromVertex := self.Vertices[from]
	toVertex := self.Vertices[to]

	fromVertex.ToEdges[to] = edge
	toVertex.FromEdges[from] = edge
}

func GenerateDirectedGraphSample() *DirectedGraph {
	//   3 -> 4            10
	//  /                 /
 	// 1 -> 5        8   9    12
	//  \    \     /      \  /
	//   2 -> 6 ->7        11
	result := &DirectedGraph{ make(map[int]*Vertex) }
	result.AddVertex(1); result.AddVertex(2); result.AddVertex(3); result.AddVertex(4)
	result.AddVertex(5); result.AddVertex(6); result.AddVertex(7); result.AddVertex(8)
	result.AddVertex(9); result.AddVertex(10); result.AddVertex(11); result.AddVertex(12)

	result.AddEdge(1, 2)
	result.AddEdge(1, 3)
	result.AddEdge(1, 5)
	result.AddEdge(2, 6)
	result.AddEdge(3, 4)
	result.AddEdge(5, 6)
	result.AddEdge(6, 7)
	result.AddEdge(7, 8)

	result.AddEdge(9, 10)
	result.AddEdge(9, 11)
	result.AddEdge(11, 12)

	return result
}
