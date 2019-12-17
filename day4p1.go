package main

/*
   a1
*/
import (
	"bufio"
	"fmt"
	"log"
	//        "math"
	"os"
	"strconv"
	"strings"
)

func main() {
	cntup := 0
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	log.Println("line1: ", stdin.Text())
	line := strings.Split(stdin.Text(), "-")
	min, _ := strconv.Atoi(line[0])
	max, _ := strconv.Atoi(line[1])
	log.Println(min, max)
	for k := min; k <= max; k++ {
		if righteg(k) && onedouble(k) {
			cntup++
		}
	}
	fmt.Println("Answer: ", cntup)
}

func onedouble(i int) bool {
	ia := inttoarray(i)

	if (ia[0] == ia[1]) ||
		ia[1] == ia[2] ||
		ia[2] == ia[3] ||
		ia[3] == ia[4] ||
		ia[4] == ia[5] {
		return true
	} else {
		return false
	}

}

func righteg(i int) bool {

	ia := inttoarray(i)

	if (ia[0] <= ia[1]) &&
		ia[1] <= ia[2] &&
		ia[2] <= ia[3] &&
		ia[3] <= ia[4] &&
		ia[4] <= ia[5] {
		return true
	} else {
		return false
	}

}

func inttoarray(i int) []int {

	var a []int = make([]int, 6)

	a[5] = i % 10 / 1
	log.Println("000001 ", i, a[5])

	a[4] = (i - a[5]*1) % 100 / 10
	log.Println("000010 ", i, a[4])

	a[3] = (i - a[4]*10 - a[5]*1) % 1000 / 100
	log.Println("000100 ", i, a[3])

	a[2] = (i - a[3]*100 - a[4]*10 - a[5]*1) % 10000 / 1000
	log.Println("001000 ", i, a[2])

	a[1] = (i - a[2]*1000 - a[3]*100 - a[4]*10 - a[5]*1) % 100000 / 10000
	log.Println("010000 ", i, a[1])

	a[0] = (i - a[1]*10000 - a[2]*1000 - a[3]*100 - a[4]*10 - a[5]*1) / 100000
	log.Println("100000 ", i, a[0])

	return a

}
