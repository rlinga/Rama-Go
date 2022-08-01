package main

import "fmt"

type Cow struct { }
type Bird struct { }
type Snake struct { }

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

 	var cmdType, animalName, actionType string

	m := make(map[string]string)

	for true {
		fmt.Println("Enter command>");
		fmt.Scanf("%s", &cmdType)
		fmt.Scanf("%s", &animalName)
		fmt.Scanf("%s", &actionType)

		switch(cmdType) {
			case "newanimal":
				m[animalName] = actionType
				fmt.Println("Created it!")
			case "query":
				var obj Animal
				_, known := m[animalName]
				if ( ! known ) {
					fmt.Println("Could not find given animal name")
					continue
				}

				switch(m[animalName]) {
					case "Cow" : obj = new(Cow)
					case "Bird" : obj = new (Bird)
					case "Snake" : obj = new (Snake)
					default: fmt.Println("Animal Name must be Cow or Bird or Snake")
						continue
				}

				switch(actionType) {
					case "Eat": obj.Eat()	
					case "Move": obj.Move()	
					case "Speak": obj.Speak()	
					default: fmt.Println("Action Type must be Eat / Move / Speak")
				}
		}
	}
}
