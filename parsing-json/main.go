package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name":"Clark",
			"last_name":"Kent",
			"hair_color":"black",
			"has_dog" : true

		},

		{
			"first_name":"David",
			"last_name":"Oluyale",
			"hair_color":"red",
			"has_dog" : true

		}

	]`

	//read json from a struct
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled:  %v", unmarshalled)

	//write json from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Jimmy"
	m2.LastName = "Pedro"
	m2.HairColor = "purple"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "		")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
