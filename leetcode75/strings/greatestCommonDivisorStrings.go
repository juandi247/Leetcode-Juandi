package main

func GreatesCommonDivisor(str1, str2 string) string {

	//they should be the same length and content, when they are backwars, if not, we know that the charachtesrs inside
	//are diff so we return inmediatly ""
	if str1+str2 != str2+str1 {
		return ""
	}

	//the content is the same so we can find the greates commno divisor
	result := gcd(len(str1), len(str2))
	return str1[:result]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	result := gcd(b, a%b)
	return result
}

/*
Como nos piden el greates common divisor, peus es lit usar la formula y ya

stack seria con 6 y 4 por jemeplo, ese vlaor es 2, porque 2 divide a 4 y a 6 al tiempo y es el valor mas grande

gcd(6,4)

gcd(4, 6%4) -> gcd(4, 2)


gdc(2, 4%2) -> gcd(2, 0) we reach our base case so we return a, its equal to 2


then this A, pops up and thats our result
*/
