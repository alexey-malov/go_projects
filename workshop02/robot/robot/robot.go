package robot

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP int64 = iota
	DOWN
	LEFT
	RIGHT
	NO_DIRECTION
)

type Robot struct {
	direction int64
	turnedOn  bool
	output    *bufio.Writer
}

func (r *Robot) TurnOn() {
	defer r.output.Flush()
	if !r.turnedOn {
		r.turnedOn = true
		fmt.Fprintln(r.output, "It am waiting for your commands")
	}
}

func (r *Robot) TurnOff() {
	defer r.output.Flush()
	if r.turnedOn {
		r.turnedOn = false
		r.direction = NO_DIRECTION
		fmt.Fprintln(r.output, "It is a pleasure to serve you")
	}
}

func (r *Robot) Walk(direction int64) {
	defer r.output.Flush()
	if r.turnedOn {
		r.direction = direction
		directions := make(map[int64]string)
		directions[UP] = "up"
		directions[DOWN] = "down"
		directions[LEFT] = "left"
		directions[RIGHT] = "right"
		fmt.Printf("Walking %v\n", directions[direction])
	} else {
		fmt.Fprintln(r.output, "The robot should be turned on first")
	}
}

func (r *Robot) IsOn() bool {
	return r.turnedOn
}

func (r *Robot) Direction() int64 {
	return r.direction
}

func (r *Robot) Output() *bufio.Writer {
	return r.output
}

func (r *Robot) Stop() {
	defer r.output.Flush()
	if r.turnedOn {
		if r.direction != NO_DIRECTION {
			r.direction = NO_DIRECTION
			fmt.Printf("Stopped\n")
		} else {
			fmt.Printf("I am staying still\n")
		}
	} else {
		fmt.Fprintln(r.output, "The robot should be turned on first")
	}
}

func NewRobot() *Robot {
	return &Robot{NO_DIRECTION, false, bufio.NewWriter(os.Stdout)}
}

func NewRobotWithWriter(output *bufio.Writer) *Robot {
	return &Robot{NO_DIRECTION, false, output}
}
