package main

import "fmt"

func main() {

	pocket := &PiggyBank{dollar: 1}
	fmt.Printf("original: %v\n", pocket.dollar)

	for i := 0; i < 10; i++ {
		saveOne(pocket)
	}
	fmt.Printf("after 10 times save: %v\n", pocket.dollar)

	for i := 0; i < 5; i++ {
		takeOne(pocket)
	}
	fmt.Printf("after 5 times take: %v\n", pocket.dollar)

}

type commenMethed interface {
	addOne() int
	subOne() int
}

func saveOne(c commenMethed) {
	c.addOne()
}

func takeOne(c commenMethed) {
	c.subOne()
}

type PiggyBank struct {
	dollar int
}

func (p *PiggyBank) addOne() int {
	p.dollar += 1
	return p.dollar
}

func (p *PiggyBank) subOne() int {
	p.dollar -= 1
	return p.dollar
}
