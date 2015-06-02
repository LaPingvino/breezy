package breezynlp

import (
	"fmt"
)

type BreezyNode struct {
	Index    int                    //Index for node
	Payload  string                 // The string to be stored in node
	Children []BreezyNeighborObject // Array of nodes connected to current node with costs to node
}

func (brNode BreezyNode) String() string {
	return fmt.Sprintf("%v %v %v\n", brNode.Index, brNode.Payload, brNode.Children)
}

func (brNode *BreezyNode) AddChild(newChild BreezyNeighborObject) bool {
	// Add child inserts a BreezyNeighberObject in to the Children array
	// If successful returns true else it returns false

	for i := 0; i < len(brNode.Children); i++ {
		if brNode.Children[i].Cost == newChild.Cost && brNode.Children[i].Vertex.Payload == newChild.Vertex.Payload {
			return false
		} else if brNode.Children[i].Cost > newChild.Cost && brNode.Children[i].Vertex.Payload == newChild.Vertex.Payload {
			brNode.Children[i].Cost = newChild.Cost
			return true
		}
	}
	brNode.Children = append(brNode.Children, newChild)
	return true
}

func (brNode *BreezyNode) removeChild(childToRemove BreezyNode) {
	for i := 0; i < len(brNode.Children); i++ {
		if brNode.Children[i].Vertex.Index == childToRemove.Index && brNode.Children[i].Vertex.Payload == childToRemove.Payload {
			tempArr := brNode.Children[i+1: len(brNode.Children)]
			brNode.Children = brNode.Children[0,i-1]
			append(brNode.Children, tempArr)
		}
	}
}

type BreezyNeighborObject struct {
	Vertex BreezyNode // Connecting neighbor node
	Cost   int        // Cost to go to new vertex from original
}

func (brNeighbor BreezyNeighborObject) String() string {
	return fmt.Sprintf("%v  %v", brNeighbor.Vertex.Payload, brNeighbor.Cost)
}

type BreezyGraph struct {
	BreezyADJList     []BreezyNode // Array of vertices within the graph
	NumberOfVerticies int          // number of vertices in the graph
	NumberOfEdges     int          // number of edges in the graph
}

func (brGraph BreezyGraph) String() string {
	return fmt.Sprintf("%v %v \n%v", brGraph.NumberOfVerticies, brGraph.NumberOfEdges, brGraph.BreezyADJList)
}

func (brGraph *BreezyGraph) AddVertex(newVertex BreezyNode) {
	// AddVertex inserts a new BreezyNode in to the BreezyADJList array
	//fmt.Println(newVertex)
	brGraph.BreezyADJList = append(brGraph.BreezyADJList, newVertex)
	brGraph.NumberOfVerticies++
}

func (brGraph *BreezyGraph) AddEdge(betweenVertex BreezyNode, andNeighbor BreezyNeighborObject) {
	//AddEdge inserts a link between two nodes that is not directed.
	isInGraph, neighborInGraph := false, false
	//Add link for initial direction
	for i := 0; i < len(brGraph.BreezyADJList); i++ {
		if brGraph.BreezyADJList[i].Index == betweenVertex.Index && brGraph.BreezyADJList[i].Payload == betweenVertex.Payload {
			isInGraph = true
			brGraph.BreezyADJList[i].AddChild(andNeighbor)

		}
		if brGraph.BreezyADJList[i].Index == andNeighbor.Vertex.Index && brGraph.BreezyADJList[i].Payload == andNeighbor.Vertex.Payload {
			neighborInGraph = true
			brGraph.BreezyADJList[i].AddChild(BreezyNeighborObject{betweenVertex, andNeighbor.Cost})
		}

	}
	if !isInGraph {
		brGraph.AddVertex(betweenVertex)
		brGraph.BreezyADJList[len(brGraph.BreezyADJList)].AddChild(andNeighbor)
	}
	if !neighborInGraph {
		brGraph.AddVertex(andNeighbor.Vertex)
		brGraph.BreezyADJList[len(brGraph.BreezyADJList)].AddChild(BreezyNeighborObject{betweenVertex, andNeighbor.Cost})
	}
	brGraph.NumberOfEdges++
}

func (brGraph *BreezyGraph) RemoveVertex(vertexToRemove BreezyNode) bool {
	// Make a queue to place other vertices connected to this vertex
	neighborQueue := BreezyQueue{nil, nil, 0}
	for i := 0; i < len(brGraph.BreezyADJList); i++ {
		if brGraph.BreezyADJList[i].Index == vertexToRemove.Index && brGraph.BreezyADJList[i].Payload == vertexToRemove.Payload {
			for j := 0; j < len(brGraph.BreezyADJList[i].Children); j++ {
				neighborQueue.enqueue(BreezyQueueNode{brGraph.BreezyADJList[i].Children[j].Vertex.Index, brGraph.BreezyADJList[i].Children[j].Vertex.Payload, nil})
			}
			// Remove children in queue
			return true
		}
	}
	return false
}

func (brGraph *BreezyGraph) RemoveEdge(fromVertex BreezyNode, andVertex BreezyNode) {
	for i := 0; i < len(brGraph.BreezyADJList); i++ {
		if brGraph.BreezyADJList[i].Index == fromVertex.Index && brGraph.BreezyADJList[i].Payload == fromVertex.Payload {
			// remove child from fromVertex
			fromVertex.removeChild(andVertex)
		}
		if brGraph.BreezyADJList[i].Index == andVertex.Index && brGraph.BreezyADJList[i].Payload == andVertex.Payload {
			// remove child from andvVrtex
			andVertex.removeChild(fromVertex)
		}
	}
}

type BreezyQueueNode struct {
	Index   int
	Payload string
	Next    *BreezyQueueNode
}
type BreezyQueue struct {
	First  *BreezyQueueNode
	Last   *BreezyQueueNode
	Length int
}

func (brQueue *BreezyQueue) enqueue(newNode BreezyQueueNode) {
	if brQueue.First == nil {
		brQueue.First = newNode
		brQueue.Last = brQueue.First
	} else {
		brQueue.Last.Next = newNode
		brQueue.Last = newNode
	}
	brQueue.Length++
}

func (brQueue *BreezyQueue) dequeue() BreezyQueueNode {
	if brQueue.First != nil {
		returnNode := brQueue.First
		brQueue.First = brQueue.First.Next
		return returnNode
	}
	return nil
}
