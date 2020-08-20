package main

import "fmt"

func main() {
	// identifier(arguments)
	foo()
	bar("James")	
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	s2, b := mouse("Ian", "Fleming")
	fmt.Println(s2)
	fmt.Println(b)
}

// func [r receiver] indentifier(parameters) [return(s)] { ... }
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