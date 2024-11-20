package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	group := flag.Bool("g", false, "group the input back into original form")
	groupsPerLine := flag.Int("n", 10, "number of groups per line when grouping")
	charactersPerGroup := flag.Int("c", 5, "number of characters per group when grouping")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	var input strings.Builder
	
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if !strings.ContainsRune(" \t\n\r", r) {
			input.WriteRune(r)
		}
	}

	result := input.String()

	if *group {
		var groupedResult strings.Builder
		runes := []rune(result)
		
		for i, r := range runes {
			groupedResult.WriteRune(r)
			if (i+1)%*charactersPerGroup == 0 && i != len(runes)-1 {
				groupedResult.WriteRune(' ')
				if (i+1)%(*charactersPerGroup**groupsPerLine) == 0 {
					groupedResult.WriteRune('\n')
				}
			}
		}
		result = groupedResult.String()
	}

	fmt.Println(result)
}
