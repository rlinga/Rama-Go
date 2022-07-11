package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var name, address string

	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter address:")
	fmt.Scan(&address)

	m := make(map[string]string)

	m["name"] = name
	m["address"] = address

	fmt.Println("Map:", m)

	j, _ := json.Marshal(m)

	fmt.Println("JSON:", string(j))
}
