package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Printf("Eats %s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("Moves : %s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("Makes this sound: %s\n", a.sound)
}

func main() {
	var animals = map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		var CheckName string
		var infoType string
		fmt.Printf(">") //Enter animal name
		fmt.Scan(&CheckName)
		if CheckName == "kill" {
			break
		}
		fmt.Printf(">") //Enter Information type (Please enter eat, move, or speak):
		fmt.Scan(&infoType)

		animal, present := animals[CheckName]
		if !present {
			fmt.Println("Animal Not Present !! , Enter Again")
			continue
		}

		switch infoType {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("enter vaild option")
		}
	}
}
