package main
 
import (
    "fmt"
)
 
// main will build and walk the tree
func main() {
    fmt.Println("Building Tree")
    root := buildTree()
    fmt.Println("Walking Tree")
    root.walk()
 
}
 
// Total is a fun way to total how many nodes we have
var total = 1
 
// How many children for the root to thave
const rootsChildren = 3
 
// How many children for the root's children to have
const childrenChildren = 10
 
// node is a super simple node struct that will form the tree
type node struct {
    parent   *node
    children []*node
    depth    int
}
 
// buildTree will construct the tree for walking.
func buildTree() *node {
    var root &node
    root.addChildren(rootsChildren)
    for _, child := range root.children {
        child.addChildren(childrenChildren)
    }
    return root
}
 
// addChildren is a convenience to add an arbitrary number of children
func (n *node) addChildren(count int) {
    for i := 0; i < count; i++ {
        newChild := &node{
            parent: n,
            depth:  n.depth + 1,
        }
        n.children = append(n.children, newChild)
    }
}
 
// walk is a recursive function that calls itself for every child
func (n *node) walk() {
    n.visit()
    for _, child := range n.children {
        child.walk()
    }
}
 
// visit will get called on every node in the tree.
func (n *node) visit() {
    d := "└"
    for i := 0; i <= n.depth; i++ {
        d = d + "───"
    }
    fmt.Printf("%s Visiting node with address %p and parent %p Total (%d)\n", d, n, n.parent, total)
    total = total + 1
}
