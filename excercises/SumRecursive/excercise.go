package sumrecursive

import "fmt"


func sumNumber(number int) int{
	if number == 1{
		return 1
	}

	return number + sumNumber(number - 1)
	}


func Test(){

	expectResult:= 15
	validate:=false

	if sumNumber(5) == expectResult{
		validate = true	
	}

	fmt.Println("Resultado ", validate)

}