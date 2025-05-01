package PointersGolang

import "fmt"

type Person struct {
	name string
	age  int
}


// ! when we return a type (struct) we are returning a copy of it. So we 
// are only able to acces it from the own function, or by making a copy of it on the Main
func personPrint() Person {
	coso := Person{name: "juandi", age: 24}
	fmt.Println(coso)
	return coso
}
// ? this would be saved on the STACK -> meaning it selfs dealloactes, by retiring it from the stack




// In this case when we return a pointer, we would have the same struct with the same memoruy addres
// within the fucntion and on main, 
func personPrintPointer() *Person{
	coso:=Person{name: "anredisto",age:23}
	return &coso
}

// ? this would be saved on the HEAP -> meaning more work for the garbage collector, but 
// ? could be good for goroutines