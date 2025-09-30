package main

// LA IDEA ERA HACERLO O(1)
func pairSum(head *ListNode) int {
	length := getLength(head)

	secondRoot := head

	target := (length / 2)
	fmt.Println("target", target)
	for i := 0; i < target; i++ {
		secondRoot = secondRoot.Next
	}
	headSecond := invertLinkedList(secondRoot)

	maxSum := 0

	for headSecond != nil {
		fmt.Println("el valro de lea head ahora es", head.Val, headSecond.Val)
		maxSum = max(maxSum, head.Val+headSecond.Val)
		head = head.Next
		headSecond = headSecond.Next
	}

	return maxSum
}

/*

Aproach: la gracia es hacerlo inplace. No cerar ningun array nuevo, ni nada asi, la idea es O(N), porque si o si toca recorrer todo
pero en espacio o (1), como??

Podriamos revertir la segunda mitad del array. Esto es o(n) tiempo pero espacio sigue siedno o(1),  luego lo qeu hacemos es ir comparanod
ambas heads, y listo. sria etnonces dos linked lists.

empezamos teneindo 5, 4, 2, 1
luego tendiramos dos linked lists 5,4, y luego 2,1

invertmios la segunda para equlibrar los pares.

5,4 y 1,2

y luego recoremos ambas al tiempo y fusrdamos en un valro la suma max actual


*/

func getLength(head *ListNode) int {
	count := 0
	for head != nil {
		head = head.Next
		count++
	}

	return count

}

func invertLinkedList(head *ListNode) *ListNode {

	currentNode := head
	previous := &ListNode{}

	for currentNode != nil {

		temp := currentNode.Next
		// break link
		currentNode.Next = previous
		// update the previous and current
		previous = currentNode
		currentNode = temp
	}

	return previous

}
