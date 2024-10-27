package main


// func arg(args ...int) {
// 	for _,num := range args {
//                fmt.Printf("%d ",num)
// 		}
// 		fmt.Printf("\n")
// 	fmt.Printf("%#v",args)
// }

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if err := root; err == nil {
		return 0
	}
	ans := 0
	ans = max(ans, 1+maxDepth(root.Right))
	ans = max(ans, 1+maxDepth(root.Left))
	return ans

	/* you can also do that
	l := 1+ maxDepth(root.Left)
	r := 1+ maxDepth(root.Right)
	return max(r,l)
	*/

}

// func pr() {
// 	count := 0

// 	// Define and immediately invoke the recursive function f
// 	var f func()
// 	f = func() {
// 		fmt.Println("Rec")

// 		if count == 2 {
// 			return
// 		}

// 		count++
// 		f() // Recursive call
// 	}
// 	f() // Invoke the function after defining it
// }

func searchmatrix(matrix [][]int, target int) bool {
	row, col := 0, 0

	for row < len(matrix) && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true

		} else if target > matrix[row][len(matrix[0])-1] {
			row++
		} else {
			col++
		}
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Dummy head for easier result list construction
	head := &ListNode{}
	temp := head
	carry := 0

	// Traverse both lists until both are fully processed
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		// Add values from l1 and l2 if they are not nil
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Calculate new carry and sum value for this node
		carry = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}
	return head.Next
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}

	return res

}

// valid perentheses problem
func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
		}
		if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}

func findDuplicate(nums []int) int {
	visit := make(map[int]bool)
	for _, n := range nums {
		_, ok := visit[n]
		if ok {
			return n
		}
		visit[n] = true
	}
	return -1

}

func subsets(nums []int) [][]int {
	var ans = [][]int{}
	var temp []int
	var helper func(int)

	helper = func(j int) {
		if j == len(nums) {
			ans = append(ans, temp[:])
			return
		}
		temp = append(temp, nums[j])
		helper(j + 1)
		temp = temp[:len(temp)-1]
		helper(j + 1)

	}
	helper(0)

	return ans
}

func combinationSum(candidates []int, target int) [][]int {

	var ans = [][]int{}
	var temp []int
	var helper func(i, sum int)

	helper = func(i, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, temp...))

			//another way to so the same thing
			// sub := make([]int , len(temp))
			// copy(sub , temp)
			// ans = append(ans, sub)

			return
		} else if sum > target || i == len(candidates) {
			return
		}

		helper(i+1, sum)
		temp = append(temp, candidates[i])
		helper(i, sum+candidates[i])
		temp = temp[:len(temp)-1]

	}

	helper(0, 0)
	return ans
}