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

        for v1 := 0; v1 <= 99; v1++ {
                for v2 := 0; v2 <= 99; v2++ {
                        dis := make([]int, len(line))
                        copy(dis, is)
                        processx(dis, v1, v2)
                }
        }

}

func processx(x []int, a1, a2 int) {

        //first instruction
        x[1] = a1
        x[2] = a2

        setnum := (len(x) - len(x)%4) / 4
        for i := 0; i <= setnum; i++ {
                if x[i*4] == 1 {
                        x = op1(x, 4*i)
                } else if x[i*4] == 2 {
                        x = op2(x, 4*i)
                } else if x[i*4] == 99 {
                        //              log.Println("Break val= ", x[i*4], " ", x)
                        break
                }
        }
        if x[0] == 19690720 {
                log.Println(x[0], " ", a1, " ", a2)
                fmt.Println("Answer: ", a1*100+a2)
        }
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
