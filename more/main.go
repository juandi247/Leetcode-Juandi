package main

import (
	"fmt"
	"sync"
)



// Estructura de Persona:
// Crea una estructura Persona que tenga Nombre, Edad y un método para calcular los años que faltan para su retiro (considera 65 años como la edad de retiro).
// Escribe funciones para crear, modificar y mostrar los datos de varias personas usando slices.


func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	messageChannel:=make(chan Person)
	usersMap:= make([]Person,0)


	
	go NewPerson("juandi",2000,messageChannel,&usersMap,&mutex )
	go NewPerson("valery",2002,messageChannel,&usersMap,&mutex )
	go NewPerson("pipe",1990,messageChannel,&usersMap,&mutex )
	go NewPerson("mama",2000,messageChannel,&usersMap,&mutex )
	go NewPerson("papa",2017,messageChannel,&usersMap,&mutex )
	go NewPerson("skipper",2010,messageChannel,&usersMap,&mutex )


		fmt.Println("voy a empezar", usersMap)

	miSlice := make([]int, 0)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(slice *[]int, wg *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			*slice = append(*slice, 1)
			mutex.Unlock()
		}(&miSlice, &wg)

		go func(slice *[]int, wg *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			*slice = append(*slice, 2)
			mutex.Unlock()

		}(&miSlice, &wg)
	}

	wg.Wait()

	 fmt.Println(miSlice)
		fmt.Println("voy a TERMINAR")


}

// Contador Concurrente:
// Escribe un programa que incremente un contador compartido desde múltiples goroutines usando un sync.Mutex para evitar condiciones de carrera.
// Luego, intenta hacerlo con un sync.WaitGroup para esperar a que todas las goroutines terminen.
