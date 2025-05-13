package somechatgptexercises

import "fmt"

// Crea una función que acepte un puntero a un slice y agregue elementos a ese slice.


func addNumbers(nums *[]int, values ...int) {
	*nums= append(*nums, values...)
}

func AppendToSlicePointer() {
    numbers := []int{1, 2, 3}
    addNumbers(&numbers, 4, 5, 6)
    fmt.Println(numbers) // debería imprimir [1, 2, 3, 4, 5, 6]
}
