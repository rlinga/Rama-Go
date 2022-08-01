package main
import ("fmt"
 "strings"
 "bufio"
 "os"
)

func main() {

	var  s string
	fmt.Println("Enter string")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s2 := scanner.Text()
		fmt.Println(s2)
		s = strings.ToLower(s2)

		if s[0] == 'i' && s[len(s)-1] == 'n' && strings.Contains(s,"a") {
			fmt.Println("Found")
		} else {
			fmt.Println("Not Found")
		}
	}
	
}
