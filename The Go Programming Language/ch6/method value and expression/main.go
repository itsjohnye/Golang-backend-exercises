package main

import (
	"fmt"
)

// type Rocket struct{}

// func (r *Rocket) Launch() {
// 	fmt.Println("Launching!")
// }

func main() {

	// r := new(Rocket)
	// time.AfterFunc(2*time.Second, r.Launch)
	// time.Sleep(3 * time.Second)

	p := Point{1, 2}
	q := Point{4, 6}
	o := Point{1, 2}
	w := Point{6, 5}

	path := Path{p, q, o, w}
	fmt.Printf("path: %v\n", path)
	path.TranslateBy(Point{1, 1}, true)
	fmt.Printf("after add {1,1}: %v\n", path)
	path.TranslateBy(Point{1, 2}, false)
	fmt.Printf("after sub {1,2}: %v\n", path)

}

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point //声明 变量op 是 结构Point 的方法
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		//调用path[i].Add(offest)或者是path[i].Sub(offset)
		path[i] = op(path[i], offset)

	}
}
