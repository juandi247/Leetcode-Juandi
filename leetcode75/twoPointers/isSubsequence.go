package main

func isSubsequence(s string, t string) bool 	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 && len(t) == 0 {
		return true
	} else if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}

	bigPointer, miniPointer := 0, 0

	for bigPointer < len(t) {

		if s[miniPointer] == t[bigPointer] {
			// match
			miniPointer++
		}

		bigPointer++
		if miniPointer == len(s) {
			return true
		}

	}

	return false

}

/*
aproach:

Vamos a tener dos poitneers como slidwing window.

Uno para el chiquito
Uno para el grande.

El chiquito recorre 1.
El grande empieza en 1.


Vamos moviendo asi sucesivamente.


Cada vez que haya un match, movemos +1 en el chiquito.

En el grande siempre vamos amover +1, encontranod o no.



ASi hasta encontrar primero que el puntero del chiquito supero al propio len de este, o que se haya acabado el otro.



*/
