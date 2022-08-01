package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	firstName string
	lastName string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFirst20(buf string) string {
	if len(buf)>20 {
		return string(buf[:20])
	} else {
		return buf
	}
}

func main() {
	var fileName string
	fmt.Println("Enter file name")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	friends := make([]Name, 0)
	for scanner.Scan() {
		fullName := scanner.Text()	
		nameParts := strings.Split(fullName, " ")
		fname := getFirst20(nameParts[0])
		lname := getFirst20(nameParts[1])
		p := Name{firstName: fname, lastName: lname}
		friends = append(friends, p)
	}

	for _, v := range friends {
		fmt.Println("First Name:", v.firstName, "Last Name:", v.lastName)
	}
}

