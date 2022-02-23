package main

import "fmt"

func main() {
	fmt.Println("Hello World.")

	var whatToSay string
	var i int

	whatToSay = "Jesus is Lord"
	fmt.Println(whatToSay)

	i = 10
	fmt.Println("i is set to ", i)

	//:= set the type to what is returned data type
	whatWasSaid, theOtherthingsaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherthingsaid)

}

//function always returns a string
func saySomething() (string, string) {
	return "something", "else"
}
