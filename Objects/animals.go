package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	var cow, bird, snake Animal

	cow.food, cow.locomotion, cow.noise = "grass", "walk", "moo"	
	bird.food, bird.locomotion, bird.noise = "worms", "fly", "peep"	
	snake.food, snake.locomotion, snake.noise = "mice", "slither", "hsss"	

	var name, action string
	var animal Animal
	
	for true {
		fmt.Println("Enter animal and characteristic interested>") 
		fmt.Scanf("%s", &name)
		fmt.Scanf("%s", &action)

		switch name {
			case "cow" : animal = cow
			case "bird" : animal = bird
			case "snake" : animal = snake
		}

		switch action {
			case "eat" : animal.Eat()
			case "move" : animal.Move()
			case "speak" : animal.Speak()
		}
	}
}
