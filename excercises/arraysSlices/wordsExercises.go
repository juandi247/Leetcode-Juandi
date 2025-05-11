package arraysSlices

import "fmt"

// Frecuencia de Palabras
// Escribe una función que reciba un slice de strings y devuelva un map con la frecuencia de cada palabra.
func CountWords(arr []string)map[string]int{

	countMap:= make(map[string]int)
	countMap["ar"]=2

	for index:= range(arr){
		value, exist :=countMap[arr[index]]
		if exist {
			countMap[arr[index]]= value +1
		}else{
			countMap[arr[index]] =1
		}

	}


	for key, value:= range(countMap){
		fmt.Printf("Word: %s, Count: %d \n", key, value)
	}

	return countMap
}