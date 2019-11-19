package graph

import (
	"fmt"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n1.AddNeighbor(n2)
	n1.AddNeighbor(n4)
	n2.AddNeighbor(n1)
	n2.AddNeighbor(n3)
	n3.AddNeighbor(n2)
	n3.AddNeighbor(n4)
	n4.AddNeighbor(n3)
	n4.AddNeighbor(n1)
	newNode1 := DeepCLoneNode(*n1)
	fmt.Printf("oldNodes map[1:%p 2:%p 3:%p 4:%p]\n", n1, n2, n3, n4)
	for _, val := range newNode1.Neighbors {
		fmt.Printf("%v\n", *val)
		for _, val2 := range val.Neighbors {
			fmt.Printf("%v\n", *val2)
		}
	}
}
