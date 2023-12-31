package main

import "fmt"

type Animal interface {
	FoodEaten() string
	Locomotion() string
	Sound() string
}

// Cow type
type cow struct {
	name string
}

func (c cow) FoodEaten() string {
	return "grass"
}
func (c cow) Locomotion() string {
	return "walk"
}
func (c cow) Sound() string {
	return "MOOOOOO"
}

// bird type
type bird struct {
	name string
}

func (b bird) FoodEaten() string {
	return "worms"
}
func (b bird) Locomotion() string {
	return "fly"
}
func (b bird) Sound() string {
	return "peep"
}

// snake type
type snake struct {
	name string
}

func (s snake) FoodEaten() string {
	return "mice"
}
func (s snake) Locomotion() string {
	return "slither"
}
func (s snake) Sound() string {
	return "hssssss"
}

func main() {
	animals := make(map[string]Animal)
	fmt.Printf("command: new animal -> cow,bird or snake \n         query -> eat,move or sound\n")
	for {
		var command, name, action string
		fmt.Printf(">")
		fmt.Scan(&command, &name, &action)

		switch command {
		case "newanimal":
			animalType := action
			var newAnimal Animal
			switch animalType {
			case "cow":
				newAnimal = cow{name}
			case "bird":
				newAnimal = bird{name}
			case "snake":
				newAnimal = snake{name}
			default:
				fmt.Println("Invalid animal type")
				continue
			}
			animals[name] = newAnimal
			fmt.Println("created it !!")

		case "query":
			animal, exits := animals[name]
			if !exits {
				fmt.Println("Animal not Found !")
				continue
			}
			switch action {
			case "eat":
				fmt.Println(animal.FoodEaten())
			case "move":
				fmt.Println(animal.Locomotion())
			case "sound":
				fmt.Println(animal.Sound())
			default:
				fmt.Println("Invalid action")
			}
		default:
			fmt.Println("Invalid Command")
		}
	}
}
