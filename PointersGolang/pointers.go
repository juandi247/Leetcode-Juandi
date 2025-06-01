package PointersGolang

import "fmt"

/* Here we have this function, if we call it on main

like this:

func main(){
a:=6
potencia(a)
}

Behin the scenes there are two different frames on the runtime Stack, so we
are Copying the value into potencia, from main, but the value inside main as we now,
will remain 6 instead of 36.

To modifiy or mantain this, we would need a pointer to it

*/
func Potencia(Value int) {
	Value *= Value
	fmt.Println(&Value, Value)
}

//! Key: inmutabilidy, we only have the value and modify it inside the funciotn

/*
Here we pass a pointer to the function, the pointer will allow us to use the same value
without copying the value of main, and modifiy it directly
*/

func PotenciaPointer(Value *int) {
	*Value *= *Value
	fmt.Println(Value, *Value)
}

//! Key: Effiecncy, but more care, becasue we are using a Pointer to the variable on the main funcition,
//! so it will be modified
