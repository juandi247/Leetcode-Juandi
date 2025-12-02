/* 
aproach:
just save the lengths of each string

when we decode de string, we just traverse the lengths that we saved,
and move the windows between the string to generate each inididually

*/
type Solution struct{
    lengthsArray []int
}

func (s *Solution) Encode(strs []string) string {

    s.lengthsArray= make([]int,len(strs))

    encoded:=""
    for i,v:=range strs{
        s.lengthsArray[i]=len(v)
        encoded=encoded+v
    }

  
fmt.Println("encoded: ", encoded)
return encoded
}




func (s *Solution) Decode(encoded string) []string {
answer:=make([]string, len(s.lengthsArray))

start:=0
end:=0

for i,v:=range s.lengthsArray{
    end+=v
    answer[i]=encoded[start:end]
    start=end
}
return answer
}


/* 


STEPS: 

Create a map of the lengths and the positions O(N max 100)

unify all the strings into one.


to decode it, we receive the string and traverse the key value storage


for eacha vlue there, we obtain the value on the string from starting point until the lenght, and so on




*/
