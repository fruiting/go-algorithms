package datastructs

type BinarySearchTree struct {
	Val   int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Val: 10,
		Left: &BinarySearchTree{
			Val: 5,
			Left: &BinarySearchTree{
				Val: 3,
			},
			Right: &BinarySearchTree{
				Val: 7,
			},
		},
		Right: &BinarySearchTree{
			Val: 15,
			Right: &BinarySearchTree{
				Val: 18,
			},
		},
	}
}

func rangeSumBST(root *BinarySearchTree, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
