package main

import "fmt"

func HasTripletSum(nums []int, target int) bool {
	i := 0
	for i < len(nums) {
		j := 0
		for j < len(nums) {
			k := 0
			for k < len(nums) {
				// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
				if i != j && j != k && i != k && nums[i]+nums[j]+nums[k] == target {
					return true
				}
				k++
			}
			j++
		}
		}
		i++
	}
	// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
	return false
}

func main() {
	fmt.Println(HasTripletSum([]int{1, 2, 3}, 12))
}
