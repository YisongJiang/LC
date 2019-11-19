package graph

import "fmt"

// Node in the graph
type Node struct {
	ID        int
	Neighbors []*Node
}

// NewNode 一个新的Node
func NewNode(nodeID int) *Node {
	n := new(Node)
	n.ID = nodeID
	n.Neighbors = make([]*Node, 0, 2)
	return n
}

// AddNeighbor add new neighbor
func (n *Node) AddNeighbor(nb *Node) {
	n.Neighbors = append(n.Neighbors, nb)
}

// DeepCLoneNode 113th problem
// In this function broad-first search is used
// In this case, a node cloud connect with itself
// In this case, we don't use the clone method in golang
func DeepCLoneNode(node Node) Node {
	nodeStacked := make(map[int]*Node, 30)
	stackOfNodes := make([]*Node, 0, 30)

	nodeNew := NewNode(node.ID)
	stackOfNodes = append(stackOfNodes, &node)
	nodeStacked[node.ID] = nodeNew

	for len(stackOfNodes) > 0 {
		f := stackOfNodes[0]
		stackOfNodes = stackOfNodes[1:] // pop
		for _, child := range f.Neighbors {
			if _, ok := nodeStacked[child.ID]; !ok {
				stackOfNodes = append(stackOfNodes, child) // push
				nodeStacked[child.ID] = NewNode(child.ID)
			}
			nodeStacked[f.ID].Neighbors = append(nodeStacked[f.ID].Neighbors, nodeStacked[child.ID])
		}
	}

	fmt.Printf("newNodes %v\n", nodeStacked)
	return *nodeNew
}
