package main

import (
	"encoding/json"
	"strings"
	"fmt"
)

type person struct {
	Firstname string
	Lastname string
	Age int
}

func main() {
	var pFirst person
	personString := strings.NewReader(`{"Firstname":"Ed", "Lastname":"Navi", "Age":20}`)
	json.NewDecoder(personString).Decode(&pFirst)

	fmt.Printf("%s %s %d", pFirst.Firstname, pFirst.Lastname, pFirst.Age)
}
