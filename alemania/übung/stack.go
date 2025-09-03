package main
/*
Valid Parentheses
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Return true if s is a valid string, and false otherwise.


Input: s = "([{}])"

Output: true


APROACH:

In general the aproach to this should be:
The First parenthesis or letter goes to the Stack
Then the next one is appended. If its differnet, we continue the STACK

if its the same as before, we POP or delete from the stack
If the stack is not empty, then its FALSE
if the stack is empty then its true



1. 
We need to create a map or something to validate the open and close
A Map is usefull for this so we would have key ")": "(" and this for all the 3 parenthesis

2. We are going to create our stack, it should be probalby a linkedlist really, so we can create it from scratch, or just an array for simple terms
Starting always with the first elment

3. Traverse each value of the List. 
	-Compare the value on the List with the previous value of the Stack. 
	- first it should be a valid symbol, not a letter, if its different we can discard it. 
	- It shoul be on the Map so we can validate the thing before of a letter easier.

	- Is the value on the STACK, the value of the KEY that we currently have? 
		-if yes we DELETE the previous value from the stack

		-If not we just append it

4. If the stack still has something then return FALSe 
if not return TRUE

*/

func isValid(s string) bool {
	stack:= []byte{}

	validMap:= map[byte]byte{
		')':'(',   
		']':'[',   
		'}':'{',   
	}

	for _,v:= range s{
		valueMap, extist:= (validMap[byte(v)])
		if !extist{
			stack = append(stack, byte(v))
		}else{
			if stack[len(stack)-1]==valueMap{
				stack= stack[:len(stack)-1]
			}else{
				return false
			}
		}
		

	}

	return len(stack)==0
    
}



type ListNode struct{
	val byte 
	next *ListNode
}

func (l *ListNode) remove(
)

