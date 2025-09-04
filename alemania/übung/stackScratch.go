package main
/*
Design a stack class that supports the push, pop, top, and getMin operations.

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
Each function should run in 
O
(
1
)
O(1) time.

*/
type node struct{
	val int
	next *node
}

type Stack struct{
	head *node
	minVal int
}


func minStack() *Stack{
	return &Stack{}
}

// func add

func (s *Stack) push(val int){
	NewNode:= &node{val: val}
	NewNode.next=s.head
	s.minVal= min(NewNode.val, s.minVal)	
	s.head=NewNode
}

func (s *Stack) pop(){
	s.head=s.head.next

}

func(s *Stack)top()*node{
	return s.head
}

func(s *Stack)min()int{
	return s.minVal
}



