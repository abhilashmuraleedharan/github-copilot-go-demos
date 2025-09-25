package main

import (
	"fmt"
	"math/rand"
	"time"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
// HasTripletSum checks if any three numbers in nums sum to the target.
func HasTripletSum(nums []int, target int) bool {
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == target {
					return true
				}
			}
		}
	}
	return false
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
// GenerateRandomNumbers returns a slice of n random integers between min and max.
func GenerateRandomNumbers(n int, min int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}
	return nums
}

func main() {
	nums := GenerateRandomNumbers(10, 1, 10)
	fmt.Println("Numbers:", nums)
	fmt.Println("HasTripletSum(nums, 12):", HasTripletSum(nums, 12))
}
