package main // change this to package main to be able to run main()

import (
	"strconv"
	"strings"

	"github.com/kashyapkurmilla/Go-structures/Packages/calculator"
)

func parseInput(raw string) calculator.CaloriesByElfList {
	input := calculator.CaloriesByElfList{}

	lines := strings.Split(raw, "\n")
	elfIndex := 0
	calorieCount := 0

	for _, l := range lines {
		if len(l) == 0 {
			input[calculator.Elf(elfIndex)] = calculator.Calories(calorieCount)
			calorieCount = 0
			elfIndex++
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			panic("can't parse the calories!")
		}
		calorieCount += c
	}

	input[calculator.Elf(elfIndex)] = calculator.Calories(calorieCount)

	return input
}
