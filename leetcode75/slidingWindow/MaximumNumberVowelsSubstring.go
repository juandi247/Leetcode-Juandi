package main

func maxVowels(s string, k int) int {

	currVowels := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currVowels++
		}
	}

	maxVowels := currVowels

	start := 0
	end := k

	for end <= len(s)-1 {
		if isVowel(s[start]) {
			currVowels--
		}

		if isVowel(s[end]) {
			currVowels++
		}

		if currVowels == k {
			return k
		}

		maxVowels = max(maxVowels, currVowels)
		start++
		end++
	}

	return maxVowels

}

func isVowel(letter byte) bool {
	if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
		return true
	}
	return false
}

/*
No hace falta brute force


Aproach:

REcorrer el primer substring

Luego ir moviendo por ventana. y listo

funcion helpér ara saber que es vocal o no




*/
