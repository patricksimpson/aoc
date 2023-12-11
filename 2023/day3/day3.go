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
	added bool
}

type sym struct {
	s string
	x int
	y int
	gears []int
}

func main() {
	// readFile, err := os.Open("test.txt")
	readFile, err := os.Open("input.txt")

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
					check = append(check, sym{s: item, x: i, y: row})
					if(num != "") {
						n, _ := strconv.Atoi(num)
						possible = append(possible, code{n: n, sx: i - len(num), ex: i - 1, y: row, added: false})	
						num = ""
					}
				} else {
					num =  num + item
				}
			} else {
				if(num != "") {
					n, _ := strconv.Atoi(num)
					possible = append(possible, code{n: n, sx: i - len(num), ex: i - 1, y: row, added: false})	
					num = ""
				}
			}
		}
		i := len(line)
		if(num != "") {
			n, _ := strconv.Atoi(num)
			possible = append(possible, code{n: n, sx: i - len(num), ex: i - 1, y: row, added: false})	
			num = ""
		}
		row = row + 1
	}

	value := 0
	gearValue := 0

	for i, poss := range possible {
		for x, s := range check {
			if(!poss.added) {
				if(s.y == poss.y - 1) {
					if(s.x >= poss.sx -1 && s.x <= poss.ex + 1) {
						value += (poss.n)
						poss.added = true
						possible[i].added = true
						if (s.s == "*") {
							fmt.Println("is gear", s, poss)
							check[x].gears = append(s.gears, poss.n)
						}
					}
				}
				if(s.y == poss.y) {
					if(s.x >= poss.sx -1 && s.x <= poss.ex + 1) {
						value += (poss.n)
						poss.added = true
						possible[i].added = true
						if (s.s == "*") {
							fmt.Println("is gear", s, poss)
							check[x].gears = append(s.gears, poss.n)
						}
					}
				}
				if(s.y == poss.y + 1) {
					if(s.x >= poss.sx -1 && s.x <= poss.ex + 1) {
						value += poss.n
						poss.added = true
						possible[i].added = true
						if (s.s == "*") {
							fmt.Println("is gear", s, poss)
							check[x].gears = append(s.gears, poss.n)
						}
					}
				}
			} else {
			}
		}
	}

  sum := 0 
	for _, poss := range check {
		if(len(poss.gears) > 1) {
			for _, v := range poss.gears {
				if(sum > 0) {
					sum = v * sum
				} else {
					sum = v
				}
			}
		}
		gearValue += sum
		fmt.Println("value", sum, gearValue)
		sum = 0
	}
	fmt.Println(value)
	fmt.Println(gearValue)
}
