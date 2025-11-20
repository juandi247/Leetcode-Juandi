/* 
not sure is the best code

apraoch

TRaverse the first digit string.
Recurse

Base case: when we are on the index of the digits = length -1. means that there is no digits left.
in this case we append the stirng that we have until now (that would be the entire combination)

Recursive:
Traverse the next digit (knowing that we are not at the end and we can do search[index+1]
we traverse the string of the next digit, and do the recursion.
passing the current string -> ab + (letter of the next) 


*/
import "strings"

func letterCombinations(digits string) []string {
    
    lettersMap:=buildMap()

    if len(digits)==1{
        newString, _:= lettersMap[digits[0]]
        return strings.Split(newString, "")
    }
    answer:= []string{}
    initialString, _:= lettersMap[digits[0]]

// 2,3
    for _,letter:=range initialString{
                            //   a
       recurse(0, len(digits), string(letter), digits, lettersMap, &answer)
    }

return answer
}
                // 1       2             d        23               
func recurse(recursiveIndex, lenPhone int,  currentString, digits string, letterMap map[byte]string, answer *[]string) {

    // base case. The index in which we are, is at the end.
    if recursiveIndex == lenPhone -1 {
        // newString:= currentString + letterIndex
        *answer = append(*answer, currentString)
        return
    }


    // if we are not on the end we are going to keep recursing, for EACh one of the nextLetters
    // take next letter 
    nextString, _:= letterMap[digits[recursiveIndex+1]]
    for _, nextLetter:= range nextString{
    
        recurse(recursiveIndex+1, lenPhone, currentString+string(nextLetter), digits, letterMap, answer)
    }

    return
}


func buildMap() (map[byte]string) {

    miMap:=map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    return miMap
}

/* 

Given a string with digits. 

Return all combinations that the number could represent. 


"23"


2= abc
3=def
