package main

import "fmt"

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
func HasTripletSum(nums []int, target int) bool {
       // Use three nested loops, but ensure that i < j < k to avoid duplicate and self-pairing
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

func main() {
	fmt.Println(HasTripletSum([]int{1, 2, 3}, 12))
}
