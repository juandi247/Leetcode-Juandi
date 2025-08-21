package mitrees

import "fmt"

/*

BFS -> Breath first search:QUEUE search first node on the left and right, and then children

DFS -> DEPTH -> Depth so we are using a STACK, we go children children children




We visit first Node.
Then both children of the first one.
Then the children of the other ones.

WE Follow the levels of the tree.

-> 			7
->		23         8
-> 	2     5     4     3


We visit the root. The root is the start of the QUEUE

We visit the Root and add the children to the queue.  7,23,28

The Delete the actual value (or print it, etc)

We go now to the next value of the queue (the 23 node), and then we add the values of the children to the queue
7,23,28,2,5 (children)
*/



func Breath_first_search(node *Node){

if node==nil{
	return
}

queue:= []*Node{node}

for len(queue)>0{

	current:= queue[0]
	queue= queue[1:]

	fmt.Printf("value: %v ,", current.Value)

	for _,v:= range current.Children{
		if v!=nil{
		queue=append(queue, v)
		}
	}

}

}

