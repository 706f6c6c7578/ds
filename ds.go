package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var decodeFlag = flag.Bool("d", false, "-d decode input")
var groupsPerLine = flag.Int("g", 10, "-g groups per line")
var singleLineOutput = flag.Bool("s", false, "-s single line output for decoding")

func invertSubstitution(substitution map[string]string) map[string]string {
	inverted := make(map[string]string)
	for key, value := range substitution {
		inverted[value] = key
	}
	return inverted
}

func decodeLetters(word string, substitution map[string]string) string {
	result := ""
	for i := 0; i < len(word); {
		match := false
		for key, value := range substitution {
			if strings.HasPrefix(word[i:], key) {
				result += value
				i += len(key)
				match = true
				break
			}
		}
		if !match {
			result += word[i : i+1]
			i++
		}
	}
	return result
}

func formatOutput(output string) string {
	if *decodeFlag || *singleLineOutput {
		return strings.ReplaceAll(output, "\n", "")
	}
	formatted := ""
	for i, r := range output {
		formatted += string(r)
		if (i+1)%5 == 0 && (i+1)%(*groupsPerLine*5) == 0 {
			formatted += "\n"
		} else if (i+1)%5 == 0 {
			formatted += " "
		}
	}
	return formatted
}

func scrambleSentence(sentence string, substitution map[string]string) string {
	if *decodeFlag {
		// Remove all spaces and newlines before decoding
		sentence = strings.ReplaceAll(sentence, "\n", "")
		sentence = strings.ReplaceAll(sentence, " ", "")
		return decodeLetters(sentence, substitution)
	} else {
		encoded := decodeLetters(sentence, substitution)
		return formatOutput(encoded)
	}
}

func main() {
	flag.Parse()

	substitution := map[string]string{
		"A": "8",
		"B": "40",
		"C": "41",
		"D": "0",
		"E": "1",
		"F": "42",
		"G": "43",
		"H": "44",
		"I": "2",
		"J": "45",
		"K": "46",
		"L": "47",
		"M": "48",
		"N": "3",
		"O": "49",
		"P": "50",
		"Q": "51",
		"R": "9",
		"S": "6",
		"T": "7",
		"U": "52",
		"V": "53",
		"W": "54",
		"X": "55",
		"Y": "56",
		"Z": "57",
		" ": "58",
		",": "59",
	}

	if *decodeFlag {
		substitution = invertSubstitution(substitution)
	}

	// Read the entire input
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		return
	}

	// Format the input as a single line
	input = bytes.ReplaceAll(input, []byte("\n"), []byte(""))

	// Create a scanner to read the input line by line
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		result := scrambleSentence(scanner.Text(), substitution)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

