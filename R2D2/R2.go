package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fst := scanner.Text()
	lst := strings.Split(fst," ")
	firstInt,_ := strconv.Atoi(lst[0])
	secondInt,_ := strconv.Atoi(lst[1])
	abs := secondInt-firstInt
	third := abs*2+firstInt
	fmt.Println(third)

}
