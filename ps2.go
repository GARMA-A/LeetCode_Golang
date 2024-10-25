package main


import ("fmt")


 // Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
func maxDepth(root *TreeNode) int {

    if err := root ; err == nil {
        return 0
    }
    ans := 0
    ans = max(ans , 1+maxDepth(root.Right))
    ans = max(ans , 1+maxDepth(root.Left))
    return ans

	/* you can also do that
	l := 1+ maxDepth(root.Left)
	r := 1+ maxDepth(root.Right)
	return max(r,l)
	*/

} 


func main2(){
       
fmt.Print("welcome to the ps2.go file")

}
