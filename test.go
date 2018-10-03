package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    c := 0
    tmp := x
    for tmp != 0 {
		c = c * 10 + tmp % 10
		tmp = tmp / 10
	}
    if c == x {
        return true
    }
    return false
}

func main(){
	fmt.Println(isPalindrome(121))
}