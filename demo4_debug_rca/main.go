package main

import "fmt"

// [AI GENERATED] LLM: GitHub Copilot, Mode: Documentation, Date: 2024-06-09
// HasTripletSum returns true if there exists a triplet in nums whose sum equals target.
func HasTripletSum(nums []int, target int) bool {
	i := 0
	for i < len(nums) {
		j := 0
		for j < len(nums) {
			k := 0
			for k < len(nums) {
				if nums[i]+nums[j]+nums[k] == target {
					return true
				}
				k++
			}
			j++
		}
		i++
	}
	return false
}

func main() {
	fmt.Println(HasTripletSum([]int{1, 2, 3}, 12))
}
