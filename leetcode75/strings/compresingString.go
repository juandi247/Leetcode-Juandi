/* 
aproach, just moving pointers as an sliding window.

Save the "start position" like on the FSM from C, and after that we traverse the number as a string meaning that

"12" would be now 1,2

so we can add

StartPosition='a' or letter
StartPosition+1= 1 digit
StartPosition+1= 2 digit
and then we update the start position ++
*/
import "strconv"

func compress(chars []byte) int {
    startIndex:=0
    currIndex:=0

    currLetter:=chars[0]
    countLetter:=0


    for currIndex<len(chars){
        if chars[currIndex]==currLetter{
            currIndex++
            countLetter++
            continue
        }

        // this would be the logic to save everything.
        chars[startIndex]=currLetter
        if countLetter>1{
            numberString:= strconv.Itoa(countLetter)
            for _,v:=range numberString{
                startIndex++
                chars[startIndex]=byte(v)
            }
        }
        startIndex++
        currLetter=chars[currIndex]
        countLetter=0
    }


// save the remaining because its ended
chars[startIndex]=currLetter
 if countLetter>1{
            numberString:= strconv.Itoa(countLetter)
            for _,v:=range numberString{
                startIndex++
                chars[startIndex]=byte(v)
            }
}

startIndex++
return startIndex

}


/*
We have a string.


3 requirermnts:

when we find a substring of 1 character alone we just appned it.

when we find a usbstring of 1 or more we apend the character and the count of those.

when we find that the count is bigger== 10 we append each character alone.

What does it mean that uses constant extra space.



Solution:

StartIndex:(slow)
CurrentIdex: (fast)

currentLetter=coso[0]
countLetter=0


while our currentIndex is on our limits, we are going to see

if we find a letter that is different than our current. we enter the logic of saving.

if the letter is the same, we keep going and move the current index



*/
