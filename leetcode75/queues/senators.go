/*
aproach: just create two queues.

dequeue always both values, but append it again to the queue, the senator that won or blocked the other, (with a higher value) so it is taken acount into a bigger round and can be affected again.

*/
func predictPartyVictory(senate string) string {
    
    currMax:= len(senate)-1

    senatorsR:=[]int{}
    senatorsD:=[]int{}


    for i,v:=range senate{
        if v=='R'{
            senatorsR=append(senatorsR, i)
        }else{
            senatorsD=append(senatorsD, i)
        }
    }

// we need to compare the initial values. 
// Dequeue both, 
// determine where we 
    for len(senatorsR)>0 && len(senatorsD)>0{
        currMax++
        r:= senatorsR[0]
        d:=senatorsD[0]

        // means that r won, so we append it to the end with a new round number
        if r<d{
            senatorsR= append(senatorsR, currMax)
        }else{
            senatorsD= append(senatorsD, currMax)
        }

        // dequeue both initial values
        senatorsR= senatorsR[1:]
        senatorsD= senatorsD[1:]
    }



    if len(senatorsR)==0{
        return "Dire"
    }

    return "Radiant"
}


/*
Understand the probelm. 
The senator wants to ban the one that appears first.

Meaning:
If D has their turn first, they can delete the first appearance of R.

IF R is first, they can delete the first appearnce of it. 
Will be intercalated.


We would use the index as a key, but adding more values to the queues.

Meaning that que are going to DEQUEUE both of the current pointers, but the one who wins
goes back to the queue, with an updated field, so that he can take the next round. as he started


RDD


R -> 0

D-> 1 2



We compare 0 with 1. 
R is smaller than D, meaning is more left, so he can decide first.

We dequeue both, but we add to the R queue the currMaxIndex+1.
currMaxIndex= len(senators)+1

R -> (2+1) 3

D-> 2

We compare again and we see that the D has won because he gets to vote before the other one. because its that round.
and 



 */

 
