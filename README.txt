Why Learn GO
=============

1. Garbage collection is easier. Its compiled language but still provides GC.
2. Simplied Object Oriented
3. Concurency is effecient
	A. Goroutines represent concurrent tasks
	B. Channels provide communication between tasks
	C. Select provides synch between tasks
4. C and C++ are compiled and run. Java is compiled into byte code and then interpreted at run time. Python is interpretive language. 
5. Only structs no classes. Structs will have data and methods.
6. No inheritance, no generics, no constructors
7. Workspace has 3 subdirs => src, bin, pkg. Workspace directory is defined by GOPATH. GOPATH=/Users/rlinga/go
8. package line is number one in the file. There is always a main package.
9. Building main package builds executable. Buiding non-main package builds only lib not executable.
10. Main package needs main().
	package main
	import "fmt"
	func main() {
		fmt.Printf("Hello Rama, Lets GO Ahead")
	}

11. import searches dirs specified by GOROOT and GOPATH
	GOROOT = /usr/local/go and GOPATH=/Users/rlinga/go
12. Go Tool commands
	A. go build (creates the executable for main package with same name as .go file)
	B. go fmt (Formats the sources files)
	C. go run hello.go (Builds and runs)
	D. go get package (Will download the package)
	E. go list (lists all installed packages)
	F. go test (tests go files that end with _test.go)
	
13. Vars
	A. var x int
	B. var x,y,z int

	C. Types => int, decimal, character, String
	D. var x int = 100
	E. var x = 100
	F. Uninitialzied vars have 0 value
	G. var x string // x = ""
	H. x := 100 // Both declares AND assigns the value, this can be done only inside a function
	I. int, int8, 16, 32, float, float8, 16, 32
	J. Unicode is 32-bit, UTF-8 is same as ASCII
14. Typedefs
	A. type Celsius float64
	B. type IDNum int
	C. var temp Celcius
	D. var pid IDNum
15. Codepoints are unicode charaters. They are called Rune in Go.
16. Runes offers =>
	=> Contains
	=> ToLower(), ToUpper()
	=> Replace(s, old, new, n)
	=> TrimSpace(s)
	=> Strconv provides converting from string to other types and vice versa
		=> itoa(), atoi(), 
17. Constants
	=> const x = 1.3
	=> const (y=4, z="Hi")
	=> type Grades int
		const (
			A Grades = itoa
			B
			C
			D
			E
		     ) // Starts at 1 for A and increments to make it unique. Like enum definitions
	
18. Control statements
	A. if x < 5 {
		fmt.Printf("Learn and Go Up")
		}
	B. for var x int=1; x < 10; x++ {
		fmt.Println(i)
	C. switch x {
		case 1: fmt.Println(x) // Exits after case without break
		case 2: fmt.Println(2) // No default from here. No break is needed
		default: fmt.Println("Default")
		}
	C. switch { // tag less switch
		case x>1: fmt.Println(x) // Exits after case without break
		case x<0: fmt.Println("X is negative") // No default from here. No break is needed
		default: fmt.Println("Default")
		}
19. Scan / read input
	A. fmt.Scan(&x)
BTW
====

1. Quadcore has 4 cores
2. GPUs have 1000+ cores
