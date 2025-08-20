package main

import (
	"fmt"

)

/* This is a tree representation like the linked list. We have a root node
that has its value, but instead of having a NEXT, it has an array of nodes
represnting the children so its like

			     NOde
			  /      \
(Children)  Node      Node

Height: the height of a tree is determinted by the longest branch. If we have a tree that 10 of the branches are height 3, but one
is height 100, then the height of the entire tree is going to be 100



Binary Tree: Obviously tree with just 2 children (max)
Manchmals wird benutzt nur
Node
left
Right


Binary Search Tree: tree with a strong orde

Leaf: just a node without children
*/
type Node struct {
	Value    int
	Children []*Node
	Left *Node
	Right *Node
}

// Traverse a tree (RECURSION!!!!)
/*
Preorder ->   Root - Node1 - Node1Left - Node1Right  ->>  Node2 - Node2Left - Node2Right
Preoder Inversed -> Root -> Node2 - Node2Right - Node2Left --> Node1 - node1Right - node2left
Postoder -> Most left Child - Right Child - Node - LAST is the ROOT
InOrder -> Go Left -> Node -> Right

*/

func InOrder(node *Node){
	if node==nil{
		return
	}

	InOrder(node.Left)
	fmt.Print(node.Value, ", ")
	InOrder(node.Right)
	
}
func PostOrder(node *Node){
	if node==nil{
		return
	}

	for _, v:= range node.Children{
		PostOrder(v)
	}
	fmt.Print(node.Value,",")
}

func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, ", ")
	for _, value := range node.Children {
		Preorder(value)
	}

}



func PreorderInversed(node *Node){


	if node==nil{
		return
	}
	fmt.Print(node.Value,", ")

	for i:=len(node.Children)-1; i>=0; i--{
		PreorderInversed(node.Children[i])
	}
}