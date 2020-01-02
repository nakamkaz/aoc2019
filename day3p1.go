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

type Apath struct {
	sx int
	sy int
	ex int
	ey int
}
type Epath struct {
	sx int
	sy int
	ex int
	ey int
}
type Pos struct {
	x int
	y int
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	line := strings.Split(stdin.Text(), ",")
	stdin.Scan()
	line2 := strings.Split(stdin.Text(), ",")
	ax, ay := wirepath(line)
	bx, by := wirepath(line2)
	log.Printf("Ax-By ")
	candi := make([]int,0)
	m1 := intersection(ax, by)
	for _, mx := range m1 {
		candi = append(candi,manhattan(0, 0, mx.x, mx.y))
		log.Println("Candi ", manhattan(0, 0, mx.x, mx.y))
	}

	log.Printf("Bx-Ay ")

	m2 := intersection(bx, ay)
	for _, mx := range m2 {
		candi = append(candi,manhattan(0, 0, mx.x, mx.y))
		log.Println("Candi ", manhattan(0, 0, mx.x, mx.y))
	}

	fmt.Println(min(candi))
}
func min (ia []int) int{
	m:=ia[0]
	for i:=1;i<=len(ia)-1;i++{
	if m>=ia[i]{
		m=ia[i]
	}
	}
	return m
}
func intersection(el []Epath, al []Apath) []Pos {
	pos := make([]Pos, 0)
	for _, q := range el {
		for _, k := range al {
			if (k.sx >= q.sx && k.sx <= q.ex) || (k.sx >= q.ex && k.sx <= q.sx) {
				if (q.sy >= k.sy && q.sy <= k.ey) || (q.sy >= k.ey && q.sy <= k.sy) {
					if !(q.sx == 0 && q.sy == 0 && k.sx == 0 && k.sy == 0) {
						log.Printf("pos ")
						log.Println(k.sx, q.sy)
						pos = append(pos, Pos{k.sx, q.sy})
					}
				}
			}
		}
	}
	return pos
}

func manhattan(x1, y1, x2, y2 int) int {
	log.Printf("x1=%v y1=%v x2=%v y2=%v x2-x1=%v y2-y1=%v\n", x1, y1, x2, y2, x2-x1, y2-y1)
	return int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
}

func wirepath(a []string) ([]Epath, []Apath) {
	px := 0
	py := 0
	ap := make([]Apath, 0)
	ep := make([]Epath, 0)
	for _, v := range a {
		log.Println(v, px, py)
		if strings.HasPrefix(v, "U") {
			t, _ := strconv.Atoi(strings.Replace(v, "U", "", 1))
			g := Apath{sx: px, sy: py, ex: px, ey: py + t}
			ap = append(ap, g)
			py += t
			log.Println(g)
		}

		if strings.HasPrefix(v, "D") {
			t, _ := strconv.Atoi(strings.Replace(v, "D", "", 1))
			g := Apath{sx: px, sy: py, ex: px, ey: py - t}
			ap = append(ap, g)
			py -= t
			log.Println(g)
		}

		if strings.HasPrefix(v, "R") {
			t, _ := strconv.Atoi(strings.Replace(v, "R", "", 1))
			g := Epath{sx: px, sy: py, ex: px + t, ey: py}
			ep = append(ep, g)
			px += t
			log.Println(g)
		}

		if strings.HasPrefix(v, "L") {
			t, _ := strconv.Atoi(strings.Replace(v, "L", "", 1))
			g := Epath{sx: px, sy: py, ex: px - t, ey: py}
			ep = append(ep, g)
			px -= t
			log.Println(g)
		}
	}
	return ep, ap
}
