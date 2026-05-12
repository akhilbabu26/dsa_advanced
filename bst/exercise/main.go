// package main

// import "fmt"

// type Node struct{
// 	Val int
// 	Right *Node
// 	Left *Node
// }

// func Insert(root *Node, val int)*Node{
// 	if root == nil{
// 		return &Node{Val: val}
// 	}

// 	if val < root.Val{
// 		root.Left = Insert(root.Left, val)
// 	}else{
// 		root.Right = Insert(root.Right, val)
// 	}

// 	return  root
// }

// func inorder(root *Node){
// 	if root == nil{
// 		return
// 	}

// 	inorder(root.Left)
// 	fmt.Println(root.Val)
// 	inorder(root.Right)
// }
// func preorder(root *Node){
// 	if root == nil{
// 		return
// 	}

// 	fmt.Println(root.Val)
// 	inorder(root.Left)
// 	inorder(root.Right)
// }

// func postorder(root *Node){
// 	if root == nil{
// 		return 
// 	}

// 	postorder(root.Left)
// 	postorder(root.Right)
// 	fmt.Println(root.Val)
// }

// func main(){
// 	var root *Node

// 	arr := []int{3,6,9,2,5,8,1,4,7}

// 	for _, v := range arr{
// 		root = Insert(root, v)
// 	}

// 	inorder(root)
// 	fmt.Println("--------------")
// 	preorder(root)
// 	fmt.Println("--------------")
// 	postorder(root)
// }


// package main

// import "fmt"

// type Node struct{
// 	Val int
// 	Left *Node
// 	Right *Node
// }

// func insert(root *Node, val int)*Node{
// 	if root == nil{
// 		return &Node{Val: val}
// 	}

// 	if val < root.Val{
// 		root.Left = insert(root.Left, val)
// 	}else{
// 		root.Right = insert(root.Right, val)
// 	}

// 	return root
// }

// func inorder(root *Node){
// 	if root == nil{
// 		return
// 	}

// 	inorder(root.Left)
// 	fmt.Println(root.Val)
// 	inorder(root.Right)
// }

// func preorder(root *Node){
// 	if root == nil{
// 		return
// 	}

// 	fmt.Println(root.Val)
// 	preorder(root.Left)
// 	preorder(root.Right)
// }

// func postorder(root *Node){
// 	if root == nil{
// 		return
// 	}

// 	postorder(root.Left) 
// 	postorder(root.Right) 
// 	fmt.Println(root.Val)
// }

// func main(){
// 	var root *Node
// 	arr := []int{7,4,1,2,5,8,9,6,3}

// 	for _, v := range arr{
// 		root = insert(root, v)
// 	}

// 	inorder(root)
// 	fmt.Println("-------------------")
// 	preorder(root)
// 	fmt.Println("-------------------")
// 	postorder(root)
// 	fmt.Println("-------------------")
// }

package main

import "fmt"

type Node struct{
	Val int
	Left *Node
	Right *Node
}

func Insert(root *Node, val int)*Node{
	if root == nil{
		return &Node{Val: val}
	}

	if val < root.Val{
		root.Left = Insert(root.Left, val)
	}else{
		root.Right = Insert(root.Right, val)
	}

	return root
}

func inorder(root *Node){
	if root == nil{
		return
	}

	inorder(root.Left)
	fmt.Println(root.Val)
	inorder(root.Right)
}

func main(){
	var root *Node

	arr := []int{9,1,7,3,8,2,4,6,5}

	for _, v := range arr{
		root = Insert(root, v)
	}

	inorder(root)
}