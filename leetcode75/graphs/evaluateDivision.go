/* 
what a nightmare of a probelm


*/
type node struct {
	letter string
	val    float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mapita, visited := createGraph(equations, values)
	answer := []float64{}
    found:=false
	var recursiveFn func(currLetter, targetLetter string, currMultiplication float64, idx int)


	recursiveFn = func(currLetter, targetLetter string, currMultiplication float64, idx int) {
   
		connections := mapita[currLetter]
        fmt.Println("Connections of the letter: ",currLetter)
        fmt.Println(connections)
		visited[currLetter] = true

		for _, n := range connections {
			// we found our target, we append, and return
            if found==true{
                return
            }
			if n.letter == targetLetter {
                fmt.Println("FOUND TARGET")
                fmt.Println("currMulti", currMultiplication)
                fmt.Println("val of multiplier, ", n.val)
                found=true
				answer = append(answer, currMultiplication*n.val)
				return
			}

			// base case 2: we already visited that letter, so we continue with the next ones
			val, _ := visited[n.letter]
			if val == true {
				continue
			}
            fmt.Println("we are going to recruse now with letter: ", n.letter)
            fmt.Println("currMultiplication was: ", currMultiplication)
            fmt.Println("currVal: ",n.val)
			recursiveFn(n.letter, targetLetter, currMultiplication*n.val, idx)
		}


	}




	for i, arr := range queries {
		currLetter := arr[0]
		targetLetter := arr[1]

		// case if we find a different letter
        _, existsFirst:= mapita[currLetter]
        _, existsSecond:= mapita[targetLetter]

        areValidLetters:= existsFirst && existsSecond
        
		if !areValidLetters {
			answer = append(answer, -1)
			continue
		}

        fmt.Printf("\n \n --- STARTING A NEW RECURSION ---\n \n ")
        fmt.Println("we start with letter", currLetter)
        fmt.Println("target: ", targetLetter)
        recursiveFn(currLetter, targetLetter, 1, i)
        restartVisited(visited)

        if found==false{
            answer=append(answer,-1)
        }

        found=false
	}

	fmt.Println("mapa", mapita)
	fmt.Println("answer: ", answer)
	return answer
}



func restartVisited(visitedMap map[string]bool){
    for key,_:=range visitedMap{
        visitedMap[key]=false
    }
}

func createGraph(equations [][]string, values []float64) (map[string][]node, map[string]bool) {
	mapita := make(map[string][]node)
	visited := make(map[string]bool)

	for i, arr := range equations {
		firstL := arr[0]
		secondL := arr[1]

		arrayFirst, exist := mapita[firstL]

		firstNode := node{
			letter: secondL,
			val:    values[i]}

		if exist {
			arrayFirst = append(arrayFirst, firstNode)
			mapita[firstL] = arrayFirst
		} else {
			mapita[firstL] = []node{firstNode}
		}

		arraySecond, exist2 := mapita[secondL]
		secondNode := node{
			letter: firstL,
			val:    1 / values[i]}
		if exist2 {
			arraySecond = append(arraySecond, secondNode)
            mapita[secondL]=arraySecond
		} else {
			mapita[secondL] = []node{secondNode}
		}

		visited[firstL] = false
		visited[secondL] = false
	}

	return mapita, visited
}

/*
We have pairs equations A MATRIX.

And values, an array of real numbers.


Where equations[i]-> a pair of equations, and values[i], reperent the
equations[i][0] / equations[i][1] = values[i]


APROACH:

The first thing that comes to mind is the following.


CREATE A GRAPH, with this graph we can obtain by traversing it, the division of values



We have A / B = 2.0


Meaning we have A -> B = 2.0

and we could obtain the previous path B -> A = 1/2.0 = 0.5



The same with B and C, we have B ->C = 3
C -> B = would be 1/3.0



And now we would have



    A     B     C

    all conected.

    A->B = 2
    B ->C= 3.

    If we want to search for a query where
    A/C, we just try to find the path from A to C.


    We enter our graph, check for the paths that exist on A
    for each path we do dfs. (we could do bfs, but in this case we are just fgoing to try dfs.)


    We check for that path and multply it fro the next.

    so A->B = 2.0

    then we search on the b side that has connections with
    B -> A
    B -> c

    in the traversal, we mark the path as visited already, when we go to A, we know we already came from there so
    we just dont continue there and we go to C. in this case, our current value is C, so its a base case, and we mark it as valid. append it. and thats it



    STEPS:


    Create a map of the paths.
    Saving the current path and the inverse
    We can create here the map of the flase and positive, as visited nodes.

    Traverse now every Query
    Create a Map of each letter,
    For each query, we are going to do that recursive function

*/
