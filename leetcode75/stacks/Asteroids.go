/* 
apraoch, just use a "stack" that compares the current asteroid with the head of it.
and thats it. 

*/
func asteroidCollision(asteroids []int) []int {
    stack:= []int{}
    
    for _,v:=range asteroids{
        stack=append(stack, v)
        // 2,-5, 5
        recurse(&stack)
    }


    return stack
}


func recurse(stack *[]int){

    s:= *stack
    // base cases
    length:=len(s)
    if length<=1{
        return
    }

    head:= length-1
    
    /*
    <- (if we append something here, its egal because they are not coliding)
    
    -> ->
    <- -> egal if we have a positive on the first value.
    
     */
    if !isPositive(s[head-1]) || isPositive(s[head]){
        return
    }

// Colision but base case because we eliminate both from the stack
    if s[head] + s[head-1]==0{
        s= s[0:head-1]
       *stack = s
        return  
    }


    // recurse
    // left is Positive ->
    // current is NEGATIVE <-
    // COLISION


/* COLISION RULES: 
if both have the same value, it explodes. so we do not append anything and we take out the value
*/
    
    negVal:= s[head]*-1
    currVal:= s[head-1]

// update the stack because either way we are going to remove that last element
    length--
    s=s[0:head]
    head--          

    if negVal > currVal{
        s[head]= negVal*-1
    }
        *stack=s
    recurse(stack)
}


func isPositive(val int)bool{
    if val>0{
        return true
    }
    return false
}

/* 
We have an array of asteorids sizes.

They are coming in the order of index.


Each sign, if its positive means RIGHT
negative LEFT.

Each asteorid moves the same speed.


RULES:

We compare the asteorid by position. Compare the colision between 2 and 1. ANd NOT the other way around.

When two asteorids have the same symbol, we ignore that and keep them there.


IF our left asteorid is negative. We can not do anything, it stays the same. 
Doesnt matter if the current one is positive or negative.


Aproach: 
We would have a kind of stack.
We traverse the whole asteroids.


For each asteroid we do the following:
Search in the left and see if: 

Our left is negative, we can NOT do nothing we CONTINUE and append it to the stack.

Our left is positive and current is positive, we append it to the stack.

Our left is positive and current is negative, we have a COLISION. 
Decide which one wins,
recurse.



Its going to be a recursive function, because its mainly a stack.




 BASE CASES: 

 If current value index is 0. 
 If left is positive and current is positive. (we append it to the stack)
 If the left value is negative and our current is positive. (we append it to the stack)



 If current is negatvie and left is positive. COlisiion
Decide which one wins
Recurse again









EXAMPLE:



3, 5, -6, 2, -1, 4.


3, 5, -6, 2.

3, 


-6, 2, 4.



*/
