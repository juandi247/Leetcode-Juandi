package MazeRecursion

import "fmt"

type Point struct {
	x int
	y int
}


func WalkingTheGrid(grid [][]int, current Point, end Point, seen[][]bool, path *[]Point) bool {

	/*
	1.Base Case: we validate that out of bounds
	validate that the point is not out of a row """current.x < len(grid[0])"""
	validate that the point is not out of a column  """current.y < len(grid)"""
	*/
	
	if current.x < 0 || current.x >= len(grid[0]) || current.y < 0 || current.y >= len(grid) {
		return false
	}
	

	/*
	2.Validate a wall
	*/
	if grid[current.x][current.y] == 1{
		return false
	}


	// 3. Check if we reached the end

	if current == end{
		*path = append(*path, current)
		return true
	}
	

	/* 4. Seen case, if we saw already this point, return false.
	for that we would need to create a grid, exactly the same of the size of the grid
	*/
	if seen[current.x][current.y] {
		return false
	}



	//pre 
	seen[current.x][current.y]= true
	//check after for pointer reference, because we are creating duplicates here
	*path = append(*path, current)



	//recursion case

	/*represent the movement using +1 -1, because on the grid, to go up 1 space, 
	it would be x, and y=y+1
	*/

	if WalkingTheGrid(grid, Point{x:current.x+1, y: current.y},end, seen, path){
		return true
	}

	if WalkingTheGrid(grid, Point{x:current.x-1, y: current.y},end, seen, path){
		return true
	}

	if WalkingTheGrid(grid, Point{x:current.x, y: current.y+1},end, seen, path){
		return true
	}

	if WalkingTheGrid(grid, Point{x:current.x, y: current.y-1},end, seen, path){
		return true
	}

	// Post
	*path = (*path)[:len(*path)-1]
	return false

}



func canReach(grid [][]int, start Point, end Point) bool {

	// define the seen array
	seen := make([][]bool, len(grid))

	for i := range seen{
		seen[i] = make([]bool, len(grid[0])) 
	}

	// define the path array
	var path []Point
	
	return WalkingTheGrid(grid,start,end,seen,&path)

}

func Test() {
	// Test Case 1
	grid1 := [][]int{
		{0, 0, 1},
		{1, 0, 1},
		{1, 0, 0},
	}


	start1 := Point{0, 0}
	end1 := Point{2, 2}
	fmt.Println("Test 1:", canReach(grid1, start1, end1)) // Esperado: true

	// Test Case 2
	grid2 := [][]int{
		{0, 1},
		{1, 0},
	}
	start2 := Point{0, 0}
	end2 := Point{1, 1}
	fmt.Println("Test 2:", canReach(grid2, start2, end2)) // Esperado: false

	// Test Case 3
	grid3 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	start3 := Point{0, 0}
	end3 := Point{2, 2}
	fmt.Println("Test 3:", canReach(grid3, start3, end3)) // Esperado: true
}
