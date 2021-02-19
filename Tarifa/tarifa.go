package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main() {
	lst := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lst = append(lst, scanner.Text())
	scanner.Scan()
	lst = append(lst, scanner.Text())
	end := lst[1]
	secondInt ,_ := strconv.Atoi(end)
	for i := 0; i < secondInt; i++ {
		scanner.Scan()
		lst = append(lst, scanner.Text())
	}
	firstInt,_ := strconv.Atoi(lst[0])
	temp := firstInt*(secondInt+1)
	elm := lst[2:]
	sum := 0
	for x := range elm {
		i,_ := strconv.Atoi(elm[x])
		sum += i
	}
	temp -= sum
	fmt.Println(temp)

}
