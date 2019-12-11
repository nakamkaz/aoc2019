package main

/*
   a1 a2 a3 ... aN
   b1 b2 b3 ... bN
*/
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
type MD {
        UD []int
        LR []int
}
*/
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	line := strings.Split(stdin.Text(), ",")
	stdin.Scan()
	line2 := strings.Split(stdin.Text(), ",")
	log.Println("\n")
	answer:= manhattan(x1,y1,x2,y2)
	fmt.Println("Answer: ",answer)
}
func manhattan(x1,y1,x2,y2 int) int {
	log.Printf("x1=%v y1=%v x2=%v y2=%v x2-x1=%v y2-y1=%v\n",x1,y1,x2,y2, x2-x1, y2-y1)
	return int( math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)) )
}

func p(a []string){
var xv, yv int
        for _, v := range line {
                if strings.HasPrefix(v, "U") {
                        t, _ := strconv.Atoi(strings.Replace(v, "U", "", 1))
                        log.Println("U +",t)
                        yv += t
                        log.Println("Y ",yv)
                }

                if strings.HasPrefix(v, "D") {
                        t, _ := strconv.Atoi(strings.Replace(v, "D", "", 1))
                        log.Println("D -",t)
                        yv -= t
                        log.Println("Y ",yv)
                }

                if strings.HasPrefix(v, "R") {
                        t, _ := strconv.Atoi(strings.Replace(v, "R", "", 1))
                        log.Println("\t\tR +",t)
                        xv += t
                        log.Println("\t\tX ",xv)
                }

                if strings.HasPrefix(v, "L") {
                        t, _ := strconv.Atoi(strings.Replace(v, "L", "", 1))
                        log.Println("\t\tL -",t)
                        xv -= t
                        log.Println("\t\tX ",xv)
                }
        }
}
