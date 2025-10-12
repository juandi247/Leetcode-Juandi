package main

func findCircleNum(isConnected [][]int) int {
	provinces := 0
	seenMap := make(map[int]bool)

	for i, _ := range isConnected {
		if !sawCity(seenMap, i) {
			provinces++
			dfs(isConnected, seenMap, i)
		}
	}
	return provinces
}

func sawCity(seenMap map[int]bool, city int) bool {
	exists, _ := seenMap[city]

	if exists {
		return true
	}
	return false
}

func dfs(grid [][]int, seenMap map[int]bool, currCity int) {
	//base case: if we seen the city, then we dont look it up
	if sawCity(seenMap, currCity) {
		fmt.Println("ciamos aca", currCity)
		return
	}

	//we need to mark our city, because we dont want to repeat that city on the next loop
	seenMap[currCity] = true

	//we need to evalute all the cities that are 1 inside of our current city
	for i, v := range grid[currCity] {
		if !sawCity(seenMap, i) && v == 1 {
			fmt.Println("encontramos una y vamos a ahcer dfs")
			dfs(grid, seenMap, i)
		}
	}
}

/*
Que me dan?

Me dan etnonce un graph en donde tenemos
1,0,0
0,1,1
0,1,1

Dos ciudades coenctadas. Serian 2 provincias porque estan ocnetadas

Seria algo mas de DEPTH etnocnes deberiamos hacerlo por depth.const


*/
