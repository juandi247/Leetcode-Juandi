/*
aproach:

just create a stack, when we have an operation, we just do it with the head and the head-1 and thats it

we pop both of them (we can just pop head, and exchange the new head with the result) and thats it

return the remaining element on th estack

*/
func evalRPN(tokens []string) int {

    stack:=[]int{}


    for _,v:=range tokens{

        if v=="+" || v=="*" || v=="-" || v=="/" {
            head:=len(stack)-1
            result:= executeOperation(stack, v,head)
            // pop
            stack= stack[:head]

            // swap minus one because we poped from the stack
            stack[head-1]=result
            continue
        }

        number,_:= strconv.Atoi(string(v))
        stack=append(stack, number)
    }



return stack[0]
}



func executeOperation(stack []int, operand string,length int) int{

    if len(stack)<2{
        fmt.Println("assertion this should neverf hapen")
        return 0
    }

result:=0
head:=stack[length]
prev:=stack[length-1]
fmt.Printf("we execute operation printiing \n head: %i \n prev: %i  \n", head, prev)
fmt.Println("stack looks like this", stack[:])
fmt.Println("operation: ", operand)
    switch operand{
        case "*":
        result= head*prev
        case "+":
        result= head+prev
        case "-":
        result= prev-head
        case "/":
        result= prev/head
       default:
       result=0 
    }


return result
}


/*
steps:

we create a stack

traverse tokens. O(N)


for each token that we find we match it with the operators. 
if its not an operator, then its a number and then we do the strconv.Atoi


if its a number we append it to the stack.




when we receive an operator value, we execute a function that checks, always that there are at least to numbers on the array,
and performs the operation.

    - pops the head
    - replaces the new head with the new result





 */
/* 
1 2 3 + / 
is a valid operation.


With this rule we are going to do:

we are going to append to the stack and execute the operation only if only, we find a symbol.


1,2,3+/ means that we could append to the stack

1,2,3 and then we reach a symbol. so we are going to pop the head, and replace now the head with the operation.

then on the next symbol, we now that we can use an operation on that symbol because there are at least 2 numbers in the stack

*/
