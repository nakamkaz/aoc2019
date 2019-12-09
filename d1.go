package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main(){
	stdin := bufio.NewScanner(os.Stdin)
	sum := 0
	for stdin.Scan(){
		a,_:= strconv.Atoi(stdin.Text())
		sum += calcvalue(a)
	}
	fmt.Println(sum)
}
func calcvalue(x int) int {
	return (x -x%3)/3 -2 
}
