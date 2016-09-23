package main

import (
	"fmt"
)
func main(){
	var num int
	fmt.Print("Enter number to start: \n")
	fmt.Scan(&num)
	for num != 1{
		if num > 0&& num % 2 == 0 {
			num = num / 2
			fmt.Println(num)
		} else if num > 0 && num % 2 != 0 {
			num = num * 3
			num = num + 1
			fmt.Println(num)
		}
	}
}
