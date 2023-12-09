package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// readFile, err := os.Open("test.txt")
	readFile, err := os.Open("input.txt")

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

		newStr := reduceString(str)
		fmt.Println(line, newStr)

		first := ""
		last := ""
		for _, elem := range newStr {
			if elem < '0' || elem > '9' {
			} else {
				if (first == "") {
					first = string(elem)
				}
				last = string(elem)
			}
    }
		value, _ := strconv.Atoi(first+last)
		fmt.Println("Line", line, ": ", value)
		
		total = total + value 
	}
	fmt.Println(total)
  readFile.Close()
}

func reduceString(str string) string {
	numbersRange := map[string] string {
		"one":"1",
		"two":"2",
		"six":"6",
		"four":"4", "five":"5",
		"nine":"9",
		"three":"3",
		"seven":"7",
		"eight":"8",
	}

	// min := 9999
	fmt.Println("before", str)

	for num, n:= range numbersRange {
		if(strings.Index(str, num) > -1) {
			end := len(num) -1
			str = strings.Replace(str, num, string(num[0]) + n + string(num[end]), -1)
		}
	}
	fmt.Println("after", str)
	return str
}
