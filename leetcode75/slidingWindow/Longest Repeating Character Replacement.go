/*
aproach:

Keep track of the most repeated element.

When the window size - most repeated elment is bigger than the K, means that the window is to big and its not posible, so we move the start +1, and acomodate everything again

 */
func characterReplacement(s string, k int) int {

freq:=make(map[byte]int)

mostFreq:=0

start:=0
end:=0

longest:=0


for end<len(s){

    // update our frequency
    letter:=s[end]
    val,_:=freq[letter]
    freq[letter]++


    val,_=freq[letter]
    mostFreq= max(mostFreq,val)

    // means that its all GOOD
    windowSize:= end-start+1
    if windowSize - mostFreq <= k || windowSize==val{
        longest= max(longest, windowSize)
    }else{
    // if its bigger, we need to resize or move the starting pointer +1
    startLetter:= s[start]
    val,_=freq[startLetter]
    freq[startLetter]=val-1
    mostFreq--
    start++
    }
    end++


}



return longest
}


/* 
What we know:

okay steps: 

We keep track of the count.
every time we do a +1, we keep the count of the biggest element



We evaluate always:
is our windwos size - biggest number <= than K?
if this is the case means that our window size is the biggest value

if NOT, we need to move our pointer of start to close the window

and decrease -1 on the elment. 
then compare now the  new value with the max, so if the max changes the next round goes better.

thats it



*/
