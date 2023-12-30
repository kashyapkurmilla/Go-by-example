// Write a program which prompts the user to enter a string.
//  The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
//  The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
//  The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

// Submit your source code for the program,
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string : ")
	// fmt.Scanln(&str)
	// fmt.Println(str)
	str, _ = reader.ReadString('\n')
	str = strings.TrimRight(str, "\n") // removes the new line character
	fmt.Println(str)
	fstr := strings.ToLower(str)
	// fmt.Println(fstr)
	if strings.Contains(fstr, "a") && strings.HasPrefix(fstr, "i") && strings.HasSuffix(fstr, "n") {
		fmt.Println("Found !!")
	} else {
		fmt.Println("Not Found !!")
	}

}
