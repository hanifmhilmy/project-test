package main

import (
	"fmt"
)

/* Time Complexity:  */

func main() {
	fmt.Println(findSequence([]int{20, 7, 8, 10, 2, 5, 6}, []int{7, 8, 10}))
	fmt.Println(findSequence([]int{20, 7, 8, 10, 2, 5, 6, 9}, []int{8, 7, 9, 10}))
	fmt.Println(findSequence([]int{20, 7, 8, 10, 2, 5, 6, 9}, []int{7, 10, 9}))
	fmt.Println(findSequence([]int{20, 7, 10, 9, 2, 5, 6, 9}, []int{7, 10, 9}))
	fmt.Println(findSequence([]int{20, 7, 10, 9, 2, 5, 6, 9}, []int{5, 6, 9}))
}

func findSequence(nums, seq []int) bool {
	if len(seq) < 1 {
		return false
	}

	marker := seq[0]
	// remove slice value from the start until we found first sequence
	for i := 0; i < len(nums); i++ {
		if nums[i] == marker {
			nums = nums[i:]
			break
		}
	}

	// loop the sequence and if found unmatched value in the index means it's not sequence
	for i := 0; i < len(seq); i++ {
		if nums[i] != seq[i] {
			return false
		}
	}

	return true
}
