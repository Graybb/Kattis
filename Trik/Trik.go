package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func swap(string2 string,i []int) []int {
	switch string2 {
	case "A":
		i[0], i[1] = i[1],i[0]
		break
	case "B":
		i[1], i[2] = i[2],i[1]
		break
	case "C":
		i[0], i[2] = i[2],i[0]
		break
	}
	return i
}

func main() {
	i := []int{1,0,0}
	reader := bufio.NewReader(os.Stdin)
	swapp, _ := reader.ReadString('\n')
	swappArray := strings.Split(swapp,"")
	for _, string := range swappArray {
		i = swap(string,i)
	}
	for x := 0; x< len(i) ; x++ {
		if i[x] ==1 {
			fmt.Print(x+1)
		}
	}

}
