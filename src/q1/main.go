package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
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
		ans += value
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}
	fmt.Printf("%d\n", ans)
}
