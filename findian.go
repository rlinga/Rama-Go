package main
import ("fmt"
 "strings"
)

func main() {

	var  s string
	
	fmt.Println("Enter a string")
	fmt.Scan(&s)
	s = strings.ToLower(s)

	if s[0] == 'i' && s[len(s)-1] == 'n' && strings.Contains(s,"a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
