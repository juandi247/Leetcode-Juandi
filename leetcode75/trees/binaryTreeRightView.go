package main

/*we need to return the MOST right nodes, of each level
thats why we need or its better with BFS


Steps.
1. Create a queue

2.0 Append the most right value of the queue to the answer array
2.1 Traverse the queue
2.2 add to the queue for each node, its sons. only if only they are not nill
2.3 Delete or pop the node that we are traversing so it gets dequeued
repeat repeat

3. return answer ezzz
*/
func RightViewNodes(root *TreeNode) []int {

	queue := []*TreeNode{root}
	answer := []int{}

	for len(queue) > 0 {
		mostRightNode := queue[len(queue)-1]

		answer = append(answer, queue[mostRightNode.Val])

		for _, v := range queue {
			//add the sons
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Left)
			}

			//dequeue the element
			queue = queue[1:]
		}

	}

	return answer
}
