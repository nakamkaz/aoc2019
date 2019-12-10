package main

/*
   a1
   a2
   a3
   ...
   aN
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	rsum int
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	sum := 0
	for stdin.Scan() {
		line, _ := strconv.Atoi(stdin.Text())
		log.Println("mass fuel ", line)
		_, result := calcvaluerc(line, 0)
		sum += result
	}
	fmt.Println(sum)
}

func calcvalue(x int) int {
	return (x-x%3)/3 - 2
}

func calcvaluerc(x, y int) (int, int) {
	log.Println("in ", x, "out", y)
	r := calcvalue(x)
	if r >= 0 {
		return calcvaluerc(r, r+y)
	} else {
		log.Println("Last in ", r, " out ", y)
		return y, y
	}
}
