
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

        //first instruction
        x[1]=12
        x[2]=2

        setnum := (len(x) - len(x)%4) / 4
        for i := 0; i <=setnum; i++ {
                if x[i*4] == 1 {
                        x = op1(x, 4*i)
                } else if x[i*4] == 2 {
                        x = op2(x, 4*i)
                } else if x[i*4] == 99 {
                        log.Println("Break val= ",x[i*4]," ",x)
                        break
                }
        }
        fmt.Println(x[0]," - Result ", x)
}

func op1(p []int, pos int) []int {
        first := pos + 1
        second := pos + 2
        third := pos + 3 // pointer to destaddr
        p[p[third]] = p[p[first]] + p[p[second]]
        return p
}

func op2(p []int, pos int) []int {
        first := pos + 1
        second := pos + 2
        third := pos + 3 // pointer to destaddr
        p[p[third]] = p[p[first]] * p[p[second]]
        return p
}
