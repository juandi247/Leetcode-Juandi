package main


func successfulPairs(spells []int, potions []int, success int64) []int {

    slices.Sort(potions)
    validSpells:= []int{}
    for _,v:=range spells{
        validSpells= append(validSpells, binarySearch(success, v, potions))
    }

return validSpells
}


func binarySearch(succes int64,multiplier int, potions []int) int{
    start:=0
    end:= len(potions)-1
    mid:= (start+end)/2
    for start<end{
     
        if start==3{
            fmt.Println("midpoint: ", mid)
        }
        result:= int64(potions[mid]*multiplier)
        if result>=succes{
            end= mid
        }else{
            start= mid+1
        }
        mid= (start+end)/2
    }

    if int64(potions[end]*multiplier) < succes{
        end++
    }

    fmt.Println("end", end)
    return len(potions) - end
}
