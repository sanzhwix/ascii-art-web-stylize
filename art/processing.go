package printing

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const bannerHeight = 8

func BannerLoader(path string) (map[rune][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	asciiMap := make(map[rune][]string)

	scanner.Scan()

	char := rune(32)

	for scanner.Scan() {
		var lines []string
		lines = append(lines, scanner.Text())
		for i := 1; i < bannerHeight; i++ {
			scanner.Scan()
			lines = append(lines, scanner.Text())
		}
		asciiMap[char] = lines
		char++
		scanner.Scan()
	}
	return asciiMap, nil
}

func Processing(text string) {
	charMap, err := BannerLoader("banners/standard.txt")
	if err != nil {
		return
	}
	text = strings.ReplaceAll(text, "\n", "\\n")
	inputLines := strings.Split(text, "\\n")
	words := false
	for i := 0; i < len(inputLines); i++ {
		if len(inputLines[i]) != 0 {
			words = true
			break
		}
	}
	for i, line := range inputLines {
		if line == "" {
			if i > 0 || words {
				fmt.Print("\n")
			}
			continue
		}

		outputLines := make([]string, 8)
		for i := 0; i < 8; i++ {
			var outputLine strings.Builder
			for _, r := range line {
				if art, ok := charMap[r]; ok && i < len(art) {
					outputLine.WriteString(art[i])
				}
			}
			outputLines[i] = outputLine.String()
		}

		for j, l := range outputLines {
			fmt.Print(l)
			if j < len(outputLines)-1 {
				fmt.Print("\n")
			}
		}

		if i < len(inputLines) {
			fmt.Print("\n")
		}
	}
}
