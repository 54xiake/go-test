package main

import "fmt"

type TreeNode struct {
	left  *TreeNode // 左子节点
	right *TreeNode // 右子节点
	value int       // 值
}

func maxDepth(root *TreeNode) int {
	// 如果节点为空就不再递归下探深度
	if root == nil {
		return 0
	}
	left := maxDepth(root.left)
	right := maxDepth(root.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	fmt.Println("tree height")
	for {
		go func() {
			fmt.Println("====")
		}()
	}
}
