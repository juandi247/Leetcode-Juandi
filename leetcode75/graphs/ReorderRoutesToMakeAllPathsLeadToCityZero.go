
func minReorder(n int, connections [][]int) int {

	mapita := make(map[int][]int)
	visited := make([]bool, n)
    count:=0
	for _, arr := range connections {

		// initialCity
		iCity := arr[0]
		// PointerCity
		ptrCity := arr[1]

		// save the REAL current road
		mapita[iCity] = append(mapita[iCity], ptrCity)

		// we save the FAKE road
		mapita[ptrCity] = append(mapita[ptrCity], -(iCity))
	}

    fmt.Println("MAPA: ", mapita)

	var recurse func(currVal int)

	recurse = func(currValue int) {
        fmt.Println("vamos a ver el value, ", currValue)
		cities, _ := mapita[currValue]
        visited[currValue]=true
        

        if currValue==4{

            fmt.Println("aray", cities)
        }
		for _, v := range cities {
            var absoluteVal int
          if v < 0 {
    absoluteVal = -v
} else {
    absoluteVal = v
}

            // base case
			if visited[absoluteVal] == true {
				continue
			}
			if v > 0 {
				count++
			}
			recurse(absoluteVal)
		}

	}


    recurse(0)

    return count

	/*
	   We should visit now every city, starting from CERO.

	   We are going to do the following:
	   The values from ZERO, if there is a negative ONe, means that okay, its a valid

	   BASE CASES:
	   If the value was already visited, return


	   If the value is negative, we dont update anything
	   and recurse

	   If the value is Positive, we update count++ and recurse

	*/

}




/*
APROACH:


We build a map of the ROADS, but first we would need to know which roads are real and which are NOT.

WE could store the value on a -sign for the not real. so we avoid using a struct



*/
