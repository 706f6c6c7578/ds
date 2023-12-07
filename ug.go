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
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	result := strings.Join(words, "")

	if *group {
		var groupedResult string
		for i, rune := range result {
			groupedResult += string(rune)
			if (i+1)%5 == 0 && i != len(result)-1 {
				groupedResult += " "
				if (i+1)%(5**groupsPerLine) == 0 {
					groupedResult += "\n"
				}
			}
		}
		result = groupedResult
	}

	fmt.Println(result)
}

