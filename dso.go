package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	decrypt := flag.Bool("d", false, "Set to true to decrypt. The operation will start from the first digit of both files.")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Please provide two filenames as arguments.")
		os.Exit(1)
	}

	// Read the files
	codeBytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keyBytes, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert the file contents to strings
	code := strings.TrimSpace(string(codeBytes))
	key := strings.TrimSpace(string(keyBytes))

	// Perform addition or subtraction based on the flag
	var result string
	if *decrypt {
		// subtraction (code - key)
		result = subtractStrings(code, key)
	} else {
		// addition (code + key)
		result = addStrings(code, key)
	}

	// Print the result
	fmt.Println(result)
}

// addStrings performs addition of two numbers represented as strings with digits
func addStrings(num1, num2 string) string {
	result := ""
	start := 0
	if len(num2) > 5 {
		result = num2[:5]
		start = 5
	}
	for i := 0; i < len(num1) && start+i < len(num2); i++ {
		digit1, _ := strconv.Atoi(string(num1[i]))
		digit2, _ := strconv.Atoi(string(num2[start+i]))
		sum := (digit1 + digit2) % 10
		result += strconv.Itoa(sum)
	}
	return result
}

// subtractStrings performs subtraction of two numbers represented as strings with digits
func subtractStrings(num1, num2 string) string {
	result := ""
	for i := 0; i < len(num1) && i < len(num2); i++ {
		digit1, _ := strconv.Atoi(string(num1[i]))
		digit2, _ := strconv.Atoi(string(num2[i]))
		diff := (digit1 - digit2 + 10) % 10
		result += strconv.Itoa(diff)
	}
	return result
}

