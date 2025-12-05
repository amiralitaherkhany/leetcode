package main

import "fmt"

func main() {
	nums := []int{10, 10, 3, 7, 6}
	result := countPartitions(nums)
	fmt.Println(result)
}

func countPartitions(nums []int) int {
	counter := 0
	for i := 1; i < len(nums); i++ {
		if (sum(nums[:i])-sum(nums[i:]))%2 == 0 {
			counter++
		}
	}
	return counter
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
