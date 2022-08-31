package awesomeProject

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	newNode := new(TreeNode)
	newNode.Val = nums[mid]

	newNode.Left = helper(nums, left, mid-1)
	newNode.Right = helper(nums, mid+1, right)
	return newNode
}
