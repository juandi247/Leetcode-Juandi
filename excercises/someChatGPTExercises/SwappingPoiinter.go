package somechatgptexercises

import "fmt"

func swap(a, b *int) {
	// derrefernce the a value, because its a pointer, so we derreference it to obtain the value
	temporalValue := *a
// asign the b value to a
	*a = *b
	*b = temporalValue
	// tu código aquí
}

func CosoMain() {
	x, y := 10, 20
	fmt.Println(&x)
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y) // debería imprimir "x: 20 y: 10"
}
