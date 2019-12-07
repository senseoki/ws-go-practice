package main

import "fmt"

func d1() {
	fmt.Print("d1() : ")
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}

}

func d2() {

	fmt.Printf("\nd2() : ")
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	fmt.Print("d3() : ")
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
