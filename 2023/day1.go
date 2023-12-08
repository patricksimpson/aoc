package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	// readFile, err := os.Open("day1-test.txt")
	readFile, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
	line := 0
	total := 0
	for fileScanner.Scan() {
		line = line + 1
		str := fileScanner.Text()

		first := ""
		last := ""
		for _, elem := range str {
			if elem < '0' || elem > '9' {
			} else {
				if (first == "") {
					first = string(elem)
				}
				last = string(elem)
			}
    }
		value, _ := strconv.Atoi(first+last)
		fmt.Println("Line: ", line, first, last, value)
		
		total = total + value 
	}
	fmt.Println(total)
  readFile.Close()
}
