package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "David",
		LastName:  "Oluyale",
	}
	myMap["me"] = me

	//SLICES
	var mySlice []string

	mySlice = append(mySlice, "David")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	//shorthand way of declaring slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	names := []string{"one", "seven", "fish", "cat"}

	log.Println(numbers[0:2])
	log.Println(names)

}
