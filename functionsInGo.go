package main

import "fmt"

type person struct{
	first	string
	last	string
	age 	int
}

type secretAgent struct{
	person
	ltk		bool
}

func main() {
	// identifier(arguments)
	foo()
	bar("James")	
	
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	
	s2, b := mouse("Ian", "Fleming")
	fmt.Println(s2)
	fmt.Println(b)
	
	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", total)

	// We can also use an unfurled slice as an argument for a func that takes variadic parameters
	xi := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	total = sum(xi...)
	fmt.Println(total)
	
	// We can even pass no arguments at all to a func that takes variadic parameters 
	total = sum()
	fmt.Println(total)

	// The keyword defer will defer the execute of the function until the moment the surrounding function returns
	defer testDefer()

	sa1 := secretAgent {
		person: person{
			"James",
			"Bond",
			32,
		},
		ltk: true,
	}
	
	sa2 := secretAgent {
		person: person{
			"Ian",
			"Fleming",
			40,
		},
		ltk: true,
	}
	sa1.speak()
	sa2.speak()
}

// func (r receiver) indentifier(parameters) return(s) { ... }
func foo() {
	fmt.Println("hello from foo")
}

// Everything in Go is passed by value
func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	s := fmt.Sprint(fn, ` `, ln, ` says "hello"`)
	b := fn == ln
	return s, b
}

// Variadic parameters are stored as a slice
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total, which is now", sum)
	}
	return sum
}

func testDefer() {
	fmt.Println("should execute last")
}

// Methods are funcs that can be attacked to any type by the use of receivers
func (s secretAgent) speak() {
	fmt.Println("hello, I'm", s.last, ",", s.first, s.last)
}