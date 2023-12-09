package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	// fmt.Println("test")
	// readFile, err := os.Open("test.txt")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
	line := 0

	var possible []int
	value := 0

	for fileScanner.Scan() {
		line = line + 1
		redMax := 0
		blueMax := 0
		greenMax := 0
		str := fileScanner.Text()
		game := strings.Split(str, ":")
		sets := strings.Split(game[1], ";")
		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, hand := range colors {
				fmt.Println(hand)
				answer := strings.Split(strings.TrimSpace(hand), " ")
				num, _ := strconv.Atoi(answer[0])
			  color := answer[1]
				fmt.Println(num, color, answer)
				if(color == "blue" && num > blueMax) {
					blueMax = num
				}
				if(color == "red" && num > redMax) {
					redMax = num
				}
				if(color == "green" && num > greenMax) {
					greenMax = num
				}
			}
		}
		power := greenMax * blueMax * redMax
		possible = append(possible, power)
		value = value + power
	}
	fmt.Println(possible)
	fmt.Println(value)
}
