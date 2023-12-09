package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type code struct {
    n int
	  sx int
	  ex int
    y int
}

type sym struct {
    s string
	  x int
    y int
}

func main() {
	readFile, err := os.Open("test.txt")
	// readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
	var possible []code
	var check []sym
	row := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		line := strings.Split(str, "")
		num := ""
		for i, item := range line {
			if (item != ".") {
				_, err := strconv.Atoi(item)
				if (err != nil) {
					check = append(check, sym{s: item, x: i , y: row})
				} else {
					num =  num + item
				}
			} else {
				if(num != "") {
					n, _ := strconv.Atoi(num)
					possible = append(possible, code{n: n, sx: i - len(num), ex: i, y: row})
					num = ""
				}
			}
		}
		row = row + 1
	}
	fmt.Println(check)
	fmt.Println(possible)

	value := 0

	for _, poss := range possible {
		for _, s := range check {
			if (((s.x - 1 >= poss.sx && s.x - 1 <= poss.ex) || (s.x + 1 >= poss.sx && s.x + 1 <= poss.ex)) &&
				((s.y -1 <= poss.y && s.y - 1 >= poss.y) || (s.y + 1 <= poss.y && s.y + 1 >= poss.y))) {
				fmt.Println("Add", poss.n, s)
				value = value + poss.n
			}
		}
	}
	fmt.Println(value)
}
