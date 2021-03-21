package main

import (
	"fmt"
	"math"
)

//Path是连接多个点的直线段
type Path []Point
type Point struct{ X, Y float64 }

//Distance方法返回路径的长度
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Path{
		{1.0, 2.0},
		{1, 3},
	}

	fmt.Println(p.Distance())

}
