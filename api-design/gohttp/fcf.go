package main

import "fmt"

// first class function

type Fn func(int, int) int

func call(sn Fn) int {
	return sn(5, 4)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fn := sum
	r1 := fn(1, 2)
	fmt.Println("fn(1, 2):", r1) // fn(1, 2): 3

	r2:= call(fn)
	fmt.Println("call(fn):", r2) // call(fn): 9

	r3 := call(sum)
	fmt.Println("call(sum):", r3) // call(sum): 9

}
