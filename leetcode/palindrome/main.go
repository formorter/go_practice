package main

// import ("fmt")

func isPalindrome(x int) bool {
	div := x

	var palindrome int
    for i := 0; div > 0; i++ {
		mod := div % 10
		div = div / 10
		palindrome = palindrome * 10 + mod
	}
	return palindrome == x
}

func main() {
	isPalindrome(12310)
}
