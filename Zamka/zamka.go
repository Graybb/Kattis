package main

import (
"bufio"
"fmt"
"os"
"strconv"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)

	L, _ := reader.ReadString('\n')
	D, _ := reader.ReadString('\n')
	X, _ := reader.ReadString('\n')
	L = strings.Trim(L, "\n")
	D = strings.Trim(D, "\n")
	X= strings.Trim(X, "\n")

	intL, _ := strconv.Atoi(L)
	intD, _ := strconv.Atoi(D)
	intX, _ := strconv.Atoi(X)
	smallest := 0
	Highest := 0
	for i := intL; i < intD+1; i++  {
		string_i := strconv.Itoa(i)
		split := strings.Split(string_i,"")
		tempDigit := 0
		for _, digit := range split {
			int_Digit,_ :=strconv.Atoi(digit)
			tempDigit += int_Digit
		}
		if tempDigit == intX {
			if  smallest == 0 {
				smallest = i
			}
			Highest = i
		}

	}
	fmt.Println(smallest)
	fmt.Print(Highest)
}
