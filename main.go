package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  int
}

type Tree struct {
	root *Node
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Numbers [Separate them into comma]: ")
	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)

	listStrings := strings.Split(text, ",")

	fmt.Printf("NUMBER INPUTS: %+v\n", listStrings)
	///------------

	pohon := Tree{}
	for _, v := range listStrings {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		pohon.InsertNode(num)
	}

	result := pohon.TraverseLeft()
	fmt.Println("ASCENDING TRAVERSE RESULT:", result)

	result = pohon.TraverseRight()
	fmt.Println("DESCENDING TRAVERSE RESULT:", result)
}

func (tree *Tree) InsertNode(value int) {
	newNode := &Node{
		value: value,
	}

	if tree.root == nil {
		tree.root = newNode
	} else {
		currentNode := tree.root
		var potentialParent *Node
		for currentNode != nil {
			potentialParent = currentNode

			if currentNode.value > value {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}

		currentNode = newNode
		currentNode.parent = potentialParent
		if currentNode.parent.value > value {
			currentNode.parent.left = newNode
		} else {
			currentNode.parent.right = newNode
		}
	}
}

func (tree *Tree) TraverseLeft() string {
	currentNode := tree.root

	result := currentNode.TraverseAscending()
	return result
}

func (node *Node) TraverseAscending() string {
	var result string

	if node.left != nil {
		nodeResult := node.left.TraverseAscending()
		result = appendString(result, nodeResult)
	}

	if node != nil {
		result = appendString(result, fmt.Sprintf("%d", node.value))
	}

	if node.right != nil {
		nodeResult := node.right.TraverseAscending()
		result = appendString(result, nodeResult)
	}

	return result
}

func (tree *Tree) TraverseRight() string {
	currentNode := tree.root

	result := currentNode.TraverseDescending()
	return result
}

func (node *Node) TraverseDescending() string {
	var result string

	if node.right != nil {
		nodeResult := node.right.TraverseDescending()
		result = appendString(result, nodeResult)
	}

	if node != nil {
		result = appendString(result, fmt.Sprintf("%d", node.value))
	}

	if node.left != nil {
		nodeResult := node.left.TraverseDescending()
		result = appendString(result, nodeResult)
	}

	return result
}

func appendString(base string, append string) string {
	if base == "" {
		return append
	}
	return fmt.Sprintf("%s | %s", base, append)
}
