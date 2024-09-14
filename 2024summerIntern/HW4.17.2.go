package main

import (
	"fmt"
)

// Node represents a node in the tree
type Node struct {
	Name     string
	Parent   *Node
	Severity int
}

// NewNode creates a new node with the given name, parent, and severity
func NewNode(name string, parent *Node, severity int) *Node {
	return &Node{
		Name:     name,
		Parent:   parent,
		Severity: severity,
	}
}

// Tree represents a tree structure
type Tree struct {
	Root *Node
}

// NewTree creates a new tree with the given root node
func NewTree(root *Node) *Tree {
	return &Tree{
		Root: root,
	}
}

// CalculateSeverity calculates the total severity of the tree
func (t *Tree) CalculateSeverity() int {
	return t.calculateNodeSeverity(t.Root)
}

// Recursive function to calculate the severity of a node and its children
func (t *Tree) calculateNodeSeverity(node *Node) int {
	severity := node.Severity
	for _, child := range t.getChildren(node) {
		severity += t.calculateNodeSeverity(child)
	}
	return severity
}

// GetChildren returns the children of a node
func (t *Tree) getChildren(node *Node) []*Node {
	children := []*Node{}
	for _, n := range nodes {
		if n.Parent == node {
			children = append(children, n)
		}
	}
	return children
}

var nodes []*Node

func main() {
	data := [][]string{
		{"a", "*", "0", "2"},
		{"b", "a", "0", "3"},
		{"b", "a", "1", "5"},
		{"a", "*", "1", "2"},
		{"e", "b", "0", "2"},
		{"f", "*", "0", "8"},
		{"c", "a", "1", "3"},
		{"d", "a", "0", "1"},
		{"d", "a", "1", "3"},
		{"f", "*", "1", "10"},
		{"g", "f", "1", "2"},
		{"h", "*", "0", "4"},
	}

	// Build the tree
	for _, item := range data {
		name := item[0]
		parentName := item[1]
		severityType := item[2]
		severityCount := item[3]

		severity := 0
		if severityType == "0" {
			severity = 5
		} else if severityType == "1" {
			severity = 2
		}

		var parent *Node
		if parentName != "*" {
			for _, n := range nodes {
				if n.Name == parentName {
					parent = n
					break
				}
			}
		}

		node := NewNode(name, parent, severity)
		nodes = append(nodes, node)
	}

	// Create trees and calculate severity for each root node
	rootNodes := make(map[*Node]bool)
	for _, node := range nodes {
		if node.Parent == nil {
			rootNodes[node] = true
		}
	}

	for rootNode := range rootNodes {
		tree := NewTree(rootNode)
		severity := tree.CalculateSeverity()
		fmt.Printf("根节点%s: 严重程度%d\n", rootNode.Name, severity)

	}
}
