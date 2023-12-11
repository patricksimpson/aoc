package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
)

func main() {
	readFile, err := os.Open("test.txt")
	// readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
	value := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		l := strings.Split(str, ":")
		gameLine := strings.Split(l[1], "|")
		answers := strings.Split(gameLine[0], " ")
		for _, hand := range answers {
			try := strings.TrimSpace(hand)
			if (try != "") {
				fmt.Println(try)
			}
		}
	}

	fmt.Println("value", value)
}
