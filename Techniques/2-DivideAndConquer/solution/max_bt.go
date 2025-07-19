package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}
	maxVal := nums[0]
	maxIndx := 0
	for i, val := range nums {
		if val > maxVal {
			maxVal = val
			maxIndx = i
		}
	}
	root := &TreeNode{
		Val:   maxVal,
		Left:  ConstructMaximumBinaryTree(nums[:maxIndx]),
		Right: ConstructMaximumBinaryTree(nums[maxIndx+1:]),
	}
	return root
}

func GetLevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			result = append(result, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
	}
	return result
}