package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

//RoverDriver 新建一台探测器
type RoverDriver struct {
	commandc chan command
}

//NewRoverDriver fn
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	start = command(0)
	stop  = command(1)
	right = command(2)
	left  = command(3)
)

//drive在goroutine中启动
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 500 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case start:
				fmt.Printf("Rover is starting at position %v\n", pos)
				speed = 1
			case stop:
				fmt.Printf("Rover stopped\n")
				speed = 0
			case right:
				fmt.Println("Trun Right")
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				fmt.Println("Trun Left")
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v; speed %d", direction, speed)
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed)) //用speed作为变量控制状态，Mul(乘以)0表示原地不动，Mul(乘以)1又启动了
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

//Left 会将探测器逆时针转90度
func (r *RoverDriver) Left() {
	r.commandc <- left
}

//Right 会将探测器顺时针转90度
func (r *RoverDriver) Right() {
	r.commandc <- right
}

//Start 探测器启动
func (r *RoverDriver) Start() {
	r.commandc <- start
}

//Stop 探测器停止
func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
