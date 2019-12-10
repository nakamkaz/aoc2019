package main

/*
   a1 a2 a3 ... aN
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// input 1,9,10,3,2,3,11,0,99,30,40,50
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	line := strings.Split(stdin.Text(), ",")
	is := make([]int, len(line))
	for i, v := range line {
		t, _ := strconv.Atoi(v)
		is[i] = t
	}
	log.Println("Data ",is)
	process(is)
}

func process(x []int) {
	setnum := (len(x) - len(x)%4) / 4
	for i := 0; i <= setnum+1; i++ {
		if x[i*4] == 1 {
			x = op1(x, 4*i)
		} else if x[i*4] == 2 {
			x = op2(x, 4*i)
		} else if x[i*4] == 99 {
			log.Println("Last ",x)
			fmt.Println("Result ", x[0])
			break
		}
	}
}

func op1(p []int, pos int) []int {
	first := pos + 1
	second := pos + 2
	third := pos + 3 // pointer to destaddr
	result := p[p[first]] + p[p[second]]
	p[p[third]] = result
	// log.Println("op1 returns ", p)
	return p
}

func op2(p []int, pos int) []int {
	first := pos + 1
	second := pos + 2
	third := pos + 3 // pointer to destaddr
	result := p[p[first]] * p[p[second]]
	p[p[third]] = result
	// log.Println("op2 returns ", p)
	return p
}
