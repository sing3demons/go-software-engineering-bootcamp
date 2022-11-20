package main

import "fmt"

type Decorator func(string) error

func Use(next Decorator) Decorator {
	fmt.Println("do something before")
	return func(s string) error {
		return next(s)
	}
}

func home(s string) error {
	fmt.Println("hello", s)
	return nil
}

// higher-order function
func main() {
	wrapped := Use(home)
	wrapped("world")
}
