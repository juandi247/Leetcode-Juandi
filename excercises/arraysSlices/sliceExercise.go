package arraysSlices

// Intersección de Slices
// Escribe una función que reciba dos slices y devuelva un slice con los elementos que están en ambos.

func IntersectTwoSlices(firstSlice []int, secondsSlice []int) []int {
	firstSlice = append(firstSlice, secondsSlice...)
	return firstSlice
}
