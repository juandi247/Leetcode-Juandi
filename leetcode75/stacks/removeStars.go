func removeStars(s string) string {

	arr := []rune{}

	for _, v := range s {
		if v != '*' {
			arr = append(arr, v)
		} else {
			arr = arr[:len(arr)-1]
		}
	}

	return string(arr)
}

/*
easiest aproach:

just create an array where we append each letter, if we find any star we pop from the array
and at the end we jsut return it as a string

*/
