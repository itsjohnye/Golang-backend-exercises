package main

import (
	"fmt"
	"math/rand"
	"time"
)

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (g gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "climbing tree"
	default:
		return "going to shit"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "apple"
	case 1:
		return "peach"
	default:
		return "banana"
	}
}

type pig struct {
	name string
}

func (g pig) String() string {
	return g.name
}

func (g pig) move() string {
	switch rand.Intn(2) {
	case 0:
		return "standing out of the house"
	default:
		return "swimming"
	}
}

func (g pig) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "rice"
	case 1:
		return "noodle"
	default:
		return "pear"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v is %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18 //设定早晚时间

func main() {
	rand.Seed(time.Now().UnixNano())

	gg := gopher{
		name: "GG",
	}

	pp := pig{
		name: "Piggy",
	}

	animals := []animal{gg, pp}

	days, hours := 3, 24 //天数和每日时间

	for day := 0; day < days; day++ {
		fmt.Printf("Day %v \n", day+1) //day+1为了消除从0开始的计数偏差
		for hour := 0; hour < hours; hour++ {
			fmt.Printf("%2d:00 ", hour)
			if hour >= sunrise && hour <= sunset {
				i := rand.Intn(len(animals))
				step(animals[i])
			} else {
				fmt.Println("The animals are sleeping.")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

}
