package main

import (
	"fmt"
)

// BTree --
type BTree struct {
	Root *BNode
}

// BNode --
type BNode struct {
	Value int
	Left  *BNode
	Right *BNode
}

// main -- package main implements Binary Tree operations
//         and solution to the House Robber problem.
func main() {
	var b BTree
	fmt.Println("***********************************")
	fmt.Println("Binary Tree")
	fmt.Println("***********************************")
	b.BinaryTree()

	fmt.Println()

	fmt.Println("***********************************")
	fmt.Println("House Robber")
	fmt.Println("***********************************")
	HouseRobber()
}

// BinaryTree --
func (bt *BTree) BinaryTree() {
	bt.Insert(10)
	bt.Insert(30)
	bt.Insert(20)
	bt.Insert(60)
	bt.Insert(40)

	fmt.Println("in-order :")
	InOrder(bt.Root)
	fmt.Println("-------------------------")

	fmt.Println("pre-order :")
	PreOrder(bt.Root)
	fmt.Println("-------------------------")

	fmt.Println("post-order :")
	PostOrder(bt.Root)
	return
}

// BTree-Insert -- add nodes at correct places in BT.
func (bt *BTree) Insert(value int) {
	// if binary tree is null, set the passed value
	// to the left of the tree.
	if bt.Root == nil {
		bt.Root = &BNode{Value: value}
		return
	}
	// if the binary tree is not null, set the passed
	// value to be inserted in child nodes.
	bt.Root.InsertInNodes(value)
}

// BNode-Insert -- Inserting the passed value in child node.
func (bn *BNode) InsertInNodes(value int) {
	// if the passed value <= parent, place it to the left
	// of the tree.
	if value <= bn.Value {
		bn.addLeft(value)
		return
	}
	// if the passed value > parent, place it to the right
	// of the tree.
	bn.addRight(value)
}

func (bn *BNode) addLeft(value int) {
	// if the left node is null, assign the passed value.
	if bn.Left == nil {
		bn.Left = &BNode{Value: value}
		return
	}
	// if left node is not null, pass the value to next node.
	bn.Left.InsertInNodes(value)
}

func (bn *BNode) addRight(value int) {
	// if the right node is null, assign the passed value.
	if bn.Right == nil {
		bn.Right = &BNode{Value: value}
		return
	}
	// if right node is not null, pass the value to next node.
	bn.Right.InsertInNodes(value)
}

// InOrder -- takes the root node of the tree as input and
// returns a list containing the In-Order Traversal of the
// given Binary Tree.
func InOrder(n *BNode) {
	if n == nil {
		return
	}

	InOrder(n.Left)
	fmt.Println(n.Value)
	InOrder(n.Right)

}

// Preorder -- takes the root node of the tree as input and
// returns a list containing the Pre-Order Traversal of the/
// given Binary Tree.
func PreOrder(n *BNode) {
	if n == nil {
		return
	}

	fmt.Println(n.Value)
	PreOrder(n.Left)
	PreOrder(n.Right)

}

// PostOrder --takes the root node of the tree as input
// and returns a list containing the Post-Order Traversal
// of the given Binary Tree.
func PostOrder(n *BNode) {
	if n == nil {
		return
	}

	PostOrder(n.Left)
	PostOrder(n.Right)
	fmt.Println(n.Value)

}

// HouseRobber --
func HouseRobber() {
	var houses []int
	oddHouses := append(houses, 3, 2, 4, 3, 2)
	fmt.Println("---odd house arrangement---")
	Robber(oddHouses)

	fmt.Println()

	houses = nil
	evenHouses := append(houses, 1, 2, 3, 4, 5, 6)
	fmt.Println("---even house arrangement---")
	Robber(evenHouses)
	fmt.Println("---------------------------")
}

// Robber --
func Robber(houses []int) {
	bountyAcquired := 0
	fmt.Println(houses)

	even := IsEven(len(houses))
	if even {
		bountyAcquired = CalulateEvenLoot(houses)
		fmt.Println("bounty acquired : $", bountyAcquired)
		return
	}

	bountyAcquired = CalulateOddLoot(houses)
	fmt.Println("bounty acquired : $", bountyAcquired)
	return

}

// Max -- returns the greatest integer value.
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// CalculateEvenLoot -- Calculates the loot acquired by the thief when the
// number of houses are even in number
func CalulateEvenLoot(houses []int) (theftBounty int) {
	for i := 0; i < len(houses); i = i + 2 {
		theftBounty += houses[i]
	}
	even := theftBounty

	// reset to 0
	theftBounty = 0
	for i := 1; i < len(houses); i = i + 2 {
		theftBounty += houses[i]
	}

	odd := theftBounty
	return Max(even, odd)
}

// CalculateOddLoot -- Calculates the loot acquired by the thief when the 
// number of houses are odd in number
func CalulateOddLoot(houses []int) (theftBounty int) {
	var slice []int
	total := len(houses)
	max := Max(houses[0], houses[total-1])
	if max == houses[0] {
		slice = houses[:total-1]
	} else {
		slice = houses[1:]
	}

	for i := 0; i < len(slice); i = i + 2 {
		theftBounty += houses[i]
	}

	even := theftBounty

	// reset to 0
	theftBounty = 0
	for i := 1; i < len(slice); i = i + 2 {
		theftBounty += houses[i]
	}

	odd := theftBounty
	return Max(even, odd)
}

// IsEven --
func IsEven(totalHouses int) bool {
	if totalHouses%2 != 0 {
		return false
	}
	return true
}
