package main

import (
	"fmt"
)

/* Time Complexity: if tree height balance, O(log n) Worst Case O(n) */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testNode := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}

	fmt.Println("Test case#1 ", searchBST(testNode, 6))
	fmt.Println("Test case#2 ", searchBST(testNode, 7))
	fmt.Println("Test case#3 ", searchBST(testNode, 8))
	fmt.Println("Test case#4 ", searchBST(testNode, 9))
	fmt.Println("Test case#5 ", searchBST(testNode, 10))
	fmt.Println("Test case#6 ", searchBST(testNode, 14))
	fmt.Println("Test case#7 ", searchBST(testNode, 3))
	fmt.Println("Test case#8 ", searchBST(testNode, 1))
}

func searchBST(root *TreeNode, val int) bool {
	for root != nil {
		if root.Val == val {
			return true
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return false
}
