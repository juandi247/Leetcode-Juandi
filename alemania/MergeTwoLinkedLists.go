package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy:= &ListNode{}
	current := dummy


	for list1!=nil && list2!=nil{
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list1 = list2.Next
		}
		current= current.Next

	}


	if list1==nil{
		// append the seconds
		current.Next= list2
	}else if list2==nil{
		// append the first list
		current.Next=list1
	}

	return dummy.Next
}

/*
We have two Lists Sorted
1,2,3,4
1,5,6,7,8

The main thing is to return a ListNode, this listnode can be the same list or a new one( this aproach with the dummy node)

Starting we should have Two Pointers
A_p -> 1
B_p -> 1

We are going to compare both values of each starting node.
If A_P is smaller than B_p then we dd to our new list that 1 from A.const

NewList-> [1]

We now move the pointer of A_p to next, so it would be 2.


Now we compare A_p again with B_p, so in this case we add the node of the B_p

NewList[1,1]

Go fuurther next on the list.






Questions or things to take into account.
If the values are equal we can take whatever we want.const

If we have 1,2, and below we have 5,6,7,7,8,9,9.const
Means that if we end comparing the first list, and now our node is negativ, means that we can append to our list, the whole
5,6,7,8,8. Etc


*/
