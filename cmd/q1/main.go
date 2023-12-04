package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <file-path>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	replaceWords := []string{"one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9"}
	replaceWords2 := []string{"eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9"}

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		replacer := strings.NewReplacer(replaceWords...)
		replacer2 := strings.NewReplacer(replaceWords2...)
		fmt.Println(line)
		line2 := reverseString(line)
		line2 = replacer2.Replace(line2)
		line2 = reverseString(line2)
		line = replacer.Replace(line)
		line = line + line2
		// Process the line as needed
		n := len(line)
		start, end := 0, n-1
		for start < n {
			if line[start] >= 48 && line[start] <= 57 {
				break
			}
			start++
		}
		for end > start {
			if line[end] >= 48 && line[end] <= 57 {
				break
			}
			end--
		}

		value, err := strconv.Atoi(string(line[start]) + string(line[end]))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(line, value)
		fmt.Println("---------------------")
		ans += value
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}
	fmt.Printf("%d\n", ans)
}
