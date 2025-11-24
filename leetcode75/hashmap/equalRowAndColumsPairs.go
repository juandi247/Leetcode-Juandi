import ("strconv"
"strings")
func equalPairs(grid [][]int) int {

columnsMap:= make(map[string]int)
    for i,_:=range grid{
        numberArray:=[]string{}
        for j,_:=range grid[0]{
            number:= grid[j][i]
            str:= strconv.Itoa(number)
            numberArray= append(numberArray, str)
        }
        columnsMap[strings.Join(numberArray, " ")]++
    }

    
count:=0
    fmt.Println("columns map", columnsMap)


    for i,_:=range grid{
        numberArray:=[]string{}
        for j,_:=range grid[0]{
            number:= grid[i][j]
            str:= strconv.Itoa(number)
            numberArray= append(numberArray, str)
        }

        numberString:= strings.Join(numberArray, " ")

        value,exists:= columnsMap[numberString]


        if exists{
            count+=value
        }

    }

    return count
}


/* 
We have a matrix.

ARPOACH:

Traverse the matrix, by COLUMNS. why? because then it would be easier the flow of code from the comparison between column and 
row


*/
