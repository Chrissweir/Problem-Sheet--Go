package main

import "fmt"

func main() {
	var primeNumber int // Store the 10000th Prime Number
	var count int // Counter
	var i int =1
	var z int
	var check bool = true //for loop bool

	for check{ // For loop exits when check set to false
		i++
		primeNumber=i // Assign i to primeNumber, when the check loop is over and the count reaches 10001,
				// this variable will contain the 10001 Prime Number
		if (i % i == 0) && (i % 1 == 0){ // Initial check for Prime Number
			if i == 2 { // Check if i is equal to 2
				count++ // Add it to the count as 2 is a prime number
			} else { // Continue to check
				for z = 2; z < i; z++ {// Check if current number can be divided by 2 or the number below it
					if i % z == 0{// If it is divisible then its not a Prime Number
						break // Try the next number
					} else if z == i-1  { //If its not divisible then its a Prime Number
						count++ //Add 1 to the count
						if count == 10001 { // When the count reaches 10001
							check = false//Exit the for loop
						}
					} else {
						continue
					}
				}
			}
		}
	}
	fmt.Print("The 10001st prime number is: ", primeNumber)// Print the 1000th Prime Number


}
