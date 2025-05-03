package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


	type Person struct {
		Name string
		Age  int
	}
	
func main() {

	// make an http call and display the json

	resp, err := http.Get("http://go.dev")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var p Person
	if e := json.NewDecoder(resp.Body).Decode(&p); e != nil {
		panic(e)
	}

	fmt.Printf("Name: %s Age: %d", p.Name, p.Age)
}
