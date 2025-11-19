
import (
    "unicode"
"strings")


/*
aproach:

create a stack, when ever we reach a closing bracket ]
we find the most proximous initial bracket [

and with that we search for the whole operation. (we search to the left of the initial bracket for the number)

POP or delete the operation from the stack (from the number to the end), and add the repeated substring. and go on





*/


func decodeString(s string) string {


    stack:=[]byte{}

    for _, v:=range s{
        stack= append(stack, byte(v))
        fmt.Println("stack", string(stack))
        if v==']'{
            fmt.Println("found a cloising")
            decodeSubstring(&stack)
        }
    }
    
    return string(stack)
}


func decodeSubstring(stack *[]byte){
    s:= *stack
    for i:=len(s)-2;i>=0; i--{
        if s[i]!='['{
            continue
        }
        fmt.Println("we found the starting cloising position: ", i)
        // open bracket position
        openP:=i
            // obtain the number even if it has two digits. check whenever we reach a letter
        startDigitP:=searchDigitIndex(s, i)
        

        // can be a 2 or 3 digit number
        number,_:= strconv.Atoi(string(s[startDigitP: openP]))
        substring:=string(s[openP+1: len(s)-1])

        // delete until the start digit so we can append now our new substring
        s= s[:startDigitP]
      
        fmt.Println("our stck looks like this now: ", string(s))
        // append the new substring
        s= append(s, generateSubstring(number, substring)...)
        break
        
    } 

    *stack = s
}


func searchDigitIndex(s []byte, lastSeenOpenBracket int) int{

// because the lastOPen bracket, we can have instead of something like 3[a2[c]] where our current bracket is the one next to the 2. so we find the next letter. But in this other case  4[2[jk]  there is no letter, so we find the [ bracket value, and we need to start on the lastOPenBracket found -1 , because if not we will return the same position as we were

     for j:=lastSeenOpenBracket-1;    j>0;    j--{
                if unicode.IsLetter(rune(s[j])) || s[j]=='[' {
                    fmt.Println("we found a letter on the stack: and index: ", string(s[j]), j)      
                    return j+1
                }
    }

    return 0

}


func generateSubstring(repeated int, substring string) []byte{
    newS:=[]byte{}

 NewSubstring := strings.Repeat(substring,repeated)
 for _,v:=range NewSubstring{
            newS= append(newS, byte(v))
}

return newS 
}
/*
We have a coded String
Decode a string

3[a]2[bc]

APRAOCH with a stack: 


We can create a stack, and when eve rwe reach a closing bracket, we find the initial bracket and its number before.
after that we ust pop from the number to the ending of the stack ,and append the value of the decoded


3[a2[c]]

would be

stack= 3[a2[c] WE reached a ] closing

we find the initial bracket, and then we find the number
2[c]. we pass that through a function, and then obtain the remaining strnig. we could even return some bytes or runes easier.








 */
