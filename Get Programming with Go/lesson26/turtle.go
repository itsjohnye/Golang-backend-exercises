package main

import (
	"fmt"
	"math/rand"
	"time"
)

type turtle struct {
	name string
	position
}

type position struct {
	x, y int
}

func (p *position) xUp() {
	p.x++
}

func (p *position) xDown() {
	p.x--
}

func (p *position) yUp() {
	p.y++
}

func (p *position) yDown() {
	p.y--
}

func main() {
	rand.Seed(time.Now().UnixNano())
	t := turtle{
		name:     "TT",
		position: position{x: 0, y: 0},
	}

	for i := 0; i < 10; i++ {
		switch rand.Intn(4) {
		case 0:
			t.xUp()
			fmt.Printf("X+1 to [%v, %v]\n", t.x, t.y)
		case 1:
			t.xDown()
			fmt.Printf("X-1 to [%v, %v]\n", t.x, t.y)
		case 2:
			t.yUp()
			fmt.Printf("Y+1 to [%v, %v]\n", t.x, t.y)
		case 3:
			t.yDown()
			fmt.Printf("Y-1 to [%v, %v]\n", t.x, t.y)
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("%v is now at the point [%v,%v]\n", t.name, t.x, t.y)
}
