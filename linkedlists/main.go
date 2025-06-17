package main



type Node struct{
	Value int
	Next *Node
}
func main() {

	head:=Node{
		Value: 3,
	}

	head.add()


}


func (n *Node) add(node *Node){
	n.Next= node
}
// ! reverse a linkedlist

func reverseList(head *Node) *Node{
	var prev *Node
	current:= head

	for current!=nil{
		temp:= current.Next
		current.Next= prev
		prev = current
		current = temp
	}
	return prev
}