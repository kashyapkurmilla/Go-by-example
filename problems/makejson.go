// Write a program which prompts the user to first enter a Name, and then enter an address.
// Your program should create a map and add the Name and address to the map using the keys “Name” and “address”, respectively. Y
// our program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var Name string
	var address string

	map1 := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Name >")
	Name, _ = reader.ReadString('\n')
	Name = strings.TrimRight(Name, "\n")

	fmt.Printf("Address >")
	address, _ = reader.ReadString('\n')
	address = strings.TrimRight(address, "\n")

	// fmt.Println(Name)
	// fmt.Println(address)
	map1["Name"] = Name
	map1["address"] = address

	jsonData, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
