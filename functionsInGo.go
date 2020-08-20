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

// Interfaces allow us to define behavior and implement polymorphism
// Any other type that has the method speak() is also of type human
// A value can be of more than one type
type human interface {
	speak()
}

// This is an empty interface
// It does not have any methods, therefore every type automatically implements this interface
type emptyInterface interface {}

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

	p1 := person{
		first: "Dr",
		last: "No",
		age: 42,
	}

	p1.speak()
	fmt.Printf("%T\n", sa1)
	fmt.Printf("%T\n", p1)

	checkType(sa1)
	checkType(sa2)
	checkType(p1)

	// Anonymous funcs
	func() {
		fmt.Println("This is an anonymous func")
	}()

	func(s string) {
		fmt.Println(s)
	}("Success!")

	// func expressions
	f1 := func(){
		fmt.Println("This is a func expression")
	}
	f1()
	
	f2 := func(i int){
		fmt.Println("The year the big brother started watching:", i)
	}
	f2(1984)

	// returnsFunc is a func that returns a nother func, which in turn returns an int
	fmt.Printf("%T\n", returnsFunc)
	fmt.Printf("%T\n", returnsFunc())

	// To get the int we need to execute the func that is returned by executing returnsFunc
	fmt.Println(returnsFunc()())
}

// func (r receiver) indentifier(parameters) return(s) { ... }
func foo() {
	fmt.Println("Hello from foo")
}

// Everything in Go is passed by value
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
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
	fmt.Println("This should execute last")
}

// Methods are funcs that can be attached to any type by the use of receivers
func (p person) speak() {
	fmt.Println("Hello, I'm", p.first, p.last)
}

func checkType(h human) {
	switch h.(type) {
		case person:
			fmt.Println("I am of type person, but I'm of also of type human.", h.(person).first)
		case secretAgent:
			fmt.Println("I am of type secretAgent, but I'm also of type human.", h.(secretAgent).ltk)
	}
}

// This func returns a func that returns an int
func returnsFunc() func() int {
	return  func() int {
		return 1984
	}
}