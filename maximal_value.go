package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max_value := nums[0]
	for i:=1; i < len(nums); i++ {
		if nums[i]>max_value{
			max_value = nums[i]
		}
	}

	fmt.Printf("Max value is: %v\n", max_value)
}
