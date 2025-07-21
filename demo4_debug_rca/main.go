package main

import "fmt"

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
        // BUG: Missing i++
    }
    return false
}

func main() {
    fmt.Println(HasTripletSum([]int{1, 2, 3}, 6))
}
