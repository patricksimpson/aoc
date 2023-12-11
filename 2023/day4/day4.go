package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
)

func main() {
	// readFile, err := os.Open("test.txt")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
	value := 0
  wins := 0
	game := 0
	for fileScanner.Scan() {
		str := fileScanner.Text()
		l := strings.Split(str, ":")
		gameLine := strings.Split(l[1], "|")
		answers := strings.Split(gameLine[0], " ")
		game += 1
		wins = 0
		for _, hand := range answers {
			try := strings.TrimSpace(hand)
			if (try != "") {
				if(hasWinner(try, strings.Split(gameLine[1], " "))) {
					if(wins > 0) {
						wins = wins + wins
					} else {
						wins = 1
					}
				}
			}
		}
		value += wins
	}
		
	fmt.Println("value", value)
	}

func hasWinner(try string, line []string) bool {
	for _, test := range line {
		num := strings.TrimSpace(test)
		if (strings.Compare(try, num) == 0) {
			return true
		}
	}
	return false
}
