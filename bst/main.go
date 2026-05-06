package main

import "fmt"

type Node struct{
	Val int
	Left *Node
	Right *Node
}

func insert(root *Node, val int)*Node{
	if root == nil{
		return &Node{Val: val}
	}

	if val < root.Val{
		root.Left = insert(root.Left, val)
	}else{
		root.Right = insert(root.Right, val)
	}

	return root
}

// SEARCH 
func search(root *Node, val int) bool {
	if root == nil {
		return false // reached end, not found
	}
	if val == root.Val {
		return true // found it!
	}
	if val < root.Val {
		return search(root.Left, val) // go left
	}
	return search(root.Right, val) // go right
}

// DELETE 
func delete(root *Node, val int) *Node {
	if root == nil {
		return nil // val not found, nothing to delete
	}

	if val < root.Val {
		root.Left = delete(root.Left, val) // go left
	} else if val > root.Val {
		root.Right = delete(root.Right, val) // go right
	} else {
		// FOUND the node to delete
		// Case 1 & 2: zero or one child
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Case 3: two children
		// Find inorder successor = smallest value in right subtree
		successor := findMin(root.Right)
		root.Val = successor.Val // copy successor value here
		root.Right = delete(root.Right, successor.Val) // delete successor
	}
	return root
}

// helper for delete — finds leftmost (smallest) node
func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func inorder(root *Node){
	if root == nil{
		return
	}

	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

func preorder(root *Node){
	if root == nil{
		return
	}

	fmt.Printf("%d ", root.Val)
    preorder(root.Left)
    preorder(root.Right)
}

func postorder(root *Node){
	if root == nil{
		return
	}

	postorder(root.Left)
	postorder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func main(){
	var root *Node

	for _, v := range []int{40, 20, 60, 10, 30, 50, 70}{
		root = insert(root, v)
	}

	fmt.Print("Inorder:   ") // 10 20 30 40 50 60 70
    inorder(root)
    fmt.Println()

    fmt.Print("Preorder:  ") // 40 20 10 30 60 50 70
    preorder(root)
    fmt.Println()

    fmt.Print("Postorder: ") // 10 30 20 50 70 60 40
    postorder(root)
    fmt.Println()

	// SEARCH
	fmt.Println("\n--- Search ---")
	fmt.Println("Search 30:", search(root, 30)) // true
	fmt.Println("Search 99:", search(root, 99)) // false

	// DELETE
	fmt.Println("\n--- Delete ---")

	// Case 1: delete leaf node (no children)
	root = delete(root, 10)
	fmt.Print("After deleting 10 (leaf):        ")
	inorder(root) // 20 30 40 50 60 70
	fmt.Println()

	// Case 2: delete node with one child
	root = delete(root, 20) // 20 now has only right child (30)
	fmt.Print("After deleting 20 (one child):   ")
	inorder(root) // 30 40 50 60 70
	fmt.Println()

	// Case 3: delete node with two children
	root = delete(root, 60) // 60 has two children: 50 and 70
	fmt.Print("After deleting 60 (two children):")
	inorder(root) // 30 40 50 70
	fmt.Println()
}