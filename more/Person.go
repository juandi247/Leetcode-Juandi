package main

import (
	"sync"
)

type Person struct {
	name     string
	YearBorn int
}

func NewPerson(name string, year int, commChannel chan Person, usersMap *[]Person, m *sync.Mutex) *Person {

	Persona := Person{
		name:     name,
		YearBorn: year,
	}

	commChannel<-Persona

	go AddToDatabase(usersMap, m, commChannel)
	return &Persona
}


func AddToDatabase(usersMap *[]Person, m *sync.Mutex, ch chan Person) {
	Persona:= <- ch
	m.Lock()
	*usersMap = append(*usersMap, Persona)
	m.Unlock()
}

func (p *Person) CalculateAge(age int) int {
	return 2025 - p.YearBorn
}

// Estructura de Persona:
// Crea una estructura Persona que tenga Nombre, Edad y un método para calcular los años que faltan para su retiro (considera 65 años como la edad de retiro).
// Escribe funciones para crear, modificar y mostrar los datos de varias personas usando slices.
