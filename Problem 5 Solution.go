package main

import (
	"fmt"
)

var fourlettercheck bool = true
var inputStr string

func main(){
	for fourlettercheck{
		fmt.Printf("Please enter 4 letters: ")
		fmt.Scan(&inputStr)

		if (len(inputStr) ==4){
			fourlettercheck =false
		}
	}
	fmt.Print("Original Word: " +inputStr+ "\nPermutations: \n", getPerms(inputStr))
}
func getPerms(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}
	current := str[0:1] // current char
	remStr := str[1:] // remaining string
	perms := getPerms(remStr) // get perms for remaining string
	allPerms := make([]string, 0) // array to hold all perms of the string based on perms of substring

	// for every perm in the perms of substring
	for _, perm := range perms {
		// add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}
	return allPerms
}

// Insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}