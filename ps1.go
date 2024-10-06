package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
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

func isPalindrome2(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		startPass := unicode.IsLetter(rune(s[l])) || unicode.IsDigit(rune(s[l]))
		endPass := unicode.IsLetter(rune(s[r])) || unicode.IsDigit(rune(s[r]))

		if !startPass {
			l++
		} else if !endPass {
			r--
		} else {
			if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func topKFrequent( nums []int , k int) []int {

	mp := make(map[int]int)
	for _,el := range nums {
        mp[el]++      
	}
	type kv struct{num int ; freq int}     
       var sl  = []kv{}
	for num,freq := range mp {
		sl = append(sl, kv{num:num , freq:freq})
	}
	sort.Slice(sl , func(i,j int) bool { return sl[i].freq > sl[j].freq})
	var res  = make([]int,0,k)
	for i :=0 ; i<k ; i++ {
		res =append(res, sl[i].num) 
	}
	return res
} 

func search(nums []int, target int) int {

    for l , r := 0 , len(nums) -1 ; l<=r ;{
        var mid = l+(r-l) /2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            r = mid-1
        } else {l = mid + 1}
    }
    return -1
    
}

func minEatingSpeed(piles []int , h int) int {
	var l , r , res int = 1 , 1e9+1 , 0
	for l<=r{
		mid := l+(r-l) / 2
		if canEat(piles , h , mid) {
                 res = mid
		   r = mid-1 
		} else {
			l = mid + 1
		}
	}
	return res

} 



func canEat(p []int,h,k int) bool {
      
	 for _,el := range p {
		h-= int(math.Ceil(float64(el) / float64(k)))
	 }
	return h>=0


}


func maxProfit(prices []int) int {
	
	front , back , mx := 0 , 1, 0
	for back < len(prices) {
		if prices[front] > prices[back] {
			front = back
		} else {
                        mx = max(mx , prices[back] - prices[front])
		}
		back++
	}
	return mx 

}

func maxProfit2(prices []int) int {
	front , back , count := 0 , 1 , 0 
       
	for back < len(prices) {
		if prices[front] > prices[back] {
			front = back
		} else if prices[front] == prices[back] {
			front++
		} else {
			count+= prices[back] - prices[front]
			front++
		}
		back++
	}

	return count


}


func main() {

fmt.Print(maxProfit([]int{7,6,4,3,1}))

}


