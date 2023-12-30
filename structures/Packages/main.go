package main // change this to package main to be able to run main()

import (
	"fmt"
	"os"

	"github.com/kashyapkurmilla/Go-structures/Packages/calculator"
)

func main() {
	list, err := os.ReadFile("./puzzle_input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	e, c := calculator.FindMostCalories(input)

	fmt.Printf("Elf %d carries %d calories\n", e, c)
}
