package main

import "fmt"

func main() {
	for i:=1; i <21; i++{
		if i % 3 == 0 {
			fmt.Printf("fizz ")
		}
		if i % 5 ==0 {
			fmt.Printf("Buzz ")
		}
		if !(i%3==0)&&!(i%5==0) {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
}
