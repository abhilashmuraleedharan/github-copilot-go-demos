package main

import (
	"fmt"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
func HasTripletSum(nums []int, target int) bool {
	for firstIdx := 0; firstIdx < len(nums); firstIdx++ {
		for secondIdx := 0; secondIdx < len(nums); secondIdx++ {
			for thirdIdx := 0; thirdIdx < len(nums); thirdIdx++ {
				if firstIdx != secondIdx && secondIdx != thirdIdx && firstIdx != thirdIdx {
					if nums[firstIdx]+nums[secondIdx]+nums[thirdIdx] == target {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(HasTripletSum([]int{1, 2, 3}, 12))
}
