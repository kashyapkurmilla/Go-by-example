package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("File not readable !!")
	}
	fmt.Println(data) //returns whole file into byte array
	fmt.Println(err)
}
