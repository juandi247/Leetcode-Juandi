package functionalprogramminggo

import "fmt"

/* this function receives 2 parameters: first a string and then
a formatter function
*/
func sayHi(name string, formatter func(string) string) {
	fmt.Println(formatter(name))
}

// we define a greeting type function called formalHi
func formalGreeting(name string) string {
	return "Good Morgin Sir/ Dr." + name
}

func mainExample() {

	// we define an anonymous function to test both of them
	hoodGreeting := func(name string) string {
		return "Sup man " + name
	}

	sayHi("Juan Diego", formalGreeting)
	sayHi("juandi", hoodGreeting)

}
