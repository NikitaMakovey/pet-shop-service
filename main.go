package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	name string
}

func (cow Cow) Eat() string {
	return "grass"
}

func (cow Cow) Move() string {
	return "walk"
}

func (cow Cow) Speak() string {
	return "moo"
}

type Bird struct {
	name string
}

func (bird Bird) Eat() string {
	return "worms"
}

func (bird Bird) Move() string {
	return "fly"
}

func (bird Bird) Speak() string {
	return "peep"
}

type Snake struct {
	name string
}

func (snake Snake) Eat() string {
	return "mice"
}

func (snake Snake) Move() string {
	return "slither"
}

func (snake Snake) Speak() string {
	return "hsss"
}

func main() {
	animals := make(map[string]Animal)
	PromptCommands(animals)
}

func PromptCommands(animals map[string]Animal) {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		requestData := strings.Split(scanner.Text(), " ")
		if len(requestData) != 3 {
			fmt.Print("> ")
			continue
		}
		command := requestData[0]
		switch command {
		case "newanimal":
			ValidateNewAnimal(animals, requestData[1], requestData[2])
		case "query":
			ValidateQuery(animals, requestData[1], requestData[2])
		}
		fmt.Print("> ")
	}
}

func ValidateNewAnimal(animals map[string]Animal, name, animalType string) {
	_, ok := animals[name]
	if !ok {
		switch animalType {
		case "cow":
			animals[name] = Cow{name}
		case "bird":
			animals[name] = Bird{name}
		case "snake":
			animals[name] = Snake{name}
		}
		fmt.Println("Created it!")
	}
}

func ValidateQuery(animals map[string]Animal, name, action string) {
	animal, ok := animals[name]
	if ok {
		switch action {
		case "eat":
			fmt.Printf("%s %s %s\n", name, "eats", animal.Eat())
		case "move":
			fmt.Printf("%s %s %s\n", name, "moves", animal.Move())
		case "speak":
			fmt.Printf("%s %s %s\n", name, "speaks", animal.Speak())
		}
	}
}
