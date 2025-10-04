package main

func canVisitAllRooms(rooms [][]int) bool {

	if len(rooms) <= 1 {
		return true
	}

	initialKeys := rooms[0]
	booleanArray := make([]bool, len(rooms))
	booleanArray[0] = true

	// 1,3
	for _, v := range initialKeys {
		recusriveFunc(v, rooms, booleanArray)

	}

	for _, v := range booleanArray {
		if !v {
			return false
		}
	}
	return true
}

func recusriveFunc(key int, rooms [][]int, booleanArray []bool) {

	if booleanArray[key] {
		return
	}

	// if it wasnt visited
	initialKeys := rooms[key]
	booleanArray[key] = true
	for _, v := range initialKeys {
		recusriveFunc(v, rooms, booleanArray)
	}
}
