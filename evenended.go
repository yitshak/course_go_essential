package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	for a:=1000; a < 9999; a++ {
		for b:=a; b <9999; b++ {
			multiplication := a*b
			mult_string := strconv.Itoa(multiplication)
			length:= len(mult_string)
			if mult_string[0]==mult_string[length-1]{
				count++;
			}
		}
	}

	fmt.Printf("There are %v round-ended number for the multiplication in checked range\n", count)
}
