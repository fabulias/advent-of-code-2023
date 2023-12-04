package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	for scanner.Scan() {
		line := scanner.Text()
		triesLine := strings.Split(line, ": ")[1]
		tries := strings.Split(triesLine, "; ")
		for _, t := range tries {
			try := strings.Split(t, ", ")
			fmt.Println(try)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}
}
