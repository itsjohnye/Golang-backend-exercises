package main

import (
	"fmt"
	"math/rand"
	"time"
)

const secondsPerDay = 86400

func main() {

	distance := 62100000 //km
	company := ""
	trip := ""

	fmt.Println("Company           Days     Trip     Price")
	fmt.Println("-----------------------------------------")

	rand.Seed(time.Now().UnixNano()) // random seed based in time
	maxSpeed := 30                   // km/s
	minSpeed := 16                   // km/s

	for i := 0; i < 10; i++ {
		//company
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		randomSpeed := rand.Intn(maxSpeed-minSpeed+1) + minSpeed //km/s
		duration := distance / randomSpeed / secondsPerDay
		price := 20 + randomSpeed

		//trip
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}
		fmt.Printf("%-18v %-4v %-10v %4v \n", company, duration, trip, price)

	}

}
