package main

import (
	 ."fmt"
	"slices"
	"strings"
	"unicode"
)

func isValidSudoku(board [][]byte) bool {
	type tuple struct{ r, c int }
	row_map := map[int][]byte{}
	col_map := map[int][]byte{}
	box_map := map[tuple][]byte{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}
			if contains(row_map[r], val) || contains(col_map[c], val) || contains(box_map[tuple{r / 3, c / 3}], val) {
				return false
			}

			row_map[r] = append(row_map[r], val)
			col_map[c] = append(col_map[c], val)
			box_map[tuple{r / 3, c / 3}] = append(box_map[tuple{r / 3, c / 3}], val)
		}
	}
	return true
}

func contains(slice []byte, val byte) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func productExceptSelf(nums []int) []int{
 // [1,2,3,4]
 size := len(nums)
 res := make([]int , size)

 prefix := 1
 postfix := 1
 for i := 0 ; i<size ; i++ {
       res[i] =  prefix 
	prefix*=nums[i]
 }
 for j := size-1 ; j>=0 ; j-- {
	res[j] *= postfix
	postfix*=nums[j]
 }

 return res
	
}


func rtrn() (int , string) {

	return 5 , "hello"
}


func isPalindrome(s string) bool {

	s1 := strings.ToLower(s)

	var cleand1 []rune

	for _,ch := range s1 {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleand1 = append(cleand1, ch)
		}
	}

	s1 = string(cleand1)

	slices.Reverse(cleand1)

	s2 := string(cleand1)

	return s1 == s2
}

func main() {
     
s := "A man, a plan, a canal: Panama"

if isPalindrome(s) {
	Println("yes")
} else {
	Println("No")
}

}


