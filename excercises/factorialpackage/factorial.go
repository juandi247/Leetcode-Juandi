package factorialpackage

// Base cases

// Factorial means number * number - 1 *number - 2, etc

func Factorial(number int) int {

	if number == 1 {
		return 1
	}
	return number * Factorial(number-1)

}
