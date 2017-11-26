package main

import (
	"fmt"
	"strings"
)

type node struct {
	children []*node
	Parent   *node
	Data     interface{}
	Level    int
}

func (n *node) AddChild(data interface{}) *node {
	newNode := createNewNode(n, n.Level+1, data)
	n.children = append(n.children, newNode)
	return newNode
}

func (n *node) AddSib(Data interface{}) *node {
	newNode := createNewNode(n.Parent, n.Level, n.Data)
	n.Parent.children = append(n.Parent.children, newNode)
	return newNode
}

func (n *node) FirstChild() *node {
	if n.HasChildren() {
		return n.children[0]
	}
	return nil
}

func (n *node) LastChild() *node {
	if n.HasChildren() {
		l := len(n.children)
		return n.children[l-1]
	}
	return nil
}

func (n *node) HasChildren() bool {
	return len(n.children) > 0
}

func (n *node) Tree() {
	if n.HasChildren() {
		for _, node := range n.children {
			fmt.Print(node.Pad())
			fmt.Println(node.Data)
			node.Tree()
		}
	}
}

func (n *node) Pad() string {
	return strings.Repeat(" ", n.Level*4)
}

func AddRoot(data interface{}) *node {
	return createNewNode(nil, 0, data)
}

func createNewNode(parent *node, level int, data interface{}) *node {
	return &node{
		children: []*node{},
		Parent:   parent,
		Level:    level,
		Data:     data,
	}
}

// Test code needs to use Goblin...

func main() {
	root := AddRoot("Root node")
	child1 := root.AddChild("Child 1")
	child2 := root.AddChild("Child 2")
	child3 := root.AddChild("Child 3")
	c3a := child3.AddChild("Child 3A")
	c3b := child3.AddChild("Child 3B")
	c3c := child3.AddChild("Child 3C")
	c1a := child1.AddChild("Child 1A")
	c1b := child1.AddChild("Child 1B")
	c1c := child1.AddChild("Child 1C")

	fmt.Println(child1.Data)
	fmt.Println(child2.Data)
	fmt.Println(child3.Data)
	fmt.Println(c1a.Data)
	fmt.Println(c1b.Data)
	fmt.Println(c1c.Data)
	fmt.Println(c3a.Data)
	fmt.Println(c3b.Data)
	fmt.Println(c3c.Data)

	fmt.Println("Iterate")
	root.Tree()
}
