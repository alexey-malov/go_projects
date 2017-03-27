package main

import (
	"bufio"
	"fmt"
	"github.com/alexey-malov/go_projects/workshop02/robot/command"
	"github.com/alexey-malov/go_projects/workshop02/robot/menu"
	"github.com/alexey-malov/go_projects/workshop02/robot/robot"
	"os"
)

func main() {
	r := robot.NewRobot()
	//robot.Stop()
	m := menu.NewMenu(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

	m.AddItem("on", "Turns the Robot on", r.TurnOn)
	m.AddItem("off", "Turns the Robot off", r.TurnOff)

	walkUp := bindWalk(r, robot.UP)
	m.AddItem("up", "Makes the Robot walk up", walkUp)
	m.AddItem("down", "Makes the Robot walk down", bindWalk(r, robot.DOWN))
	m.AddItem("left", "Makes the Robot walk left", bindWalk(r, robot.LEFT))
	walkRight := bindWalk(r, robot.RIGHT)
	m.AddItem("right", "Makes the Robot walk right", bindWalk(r, robot.RIGHT))

	m.AddItem("horse_moving", "Makes the Robot walk like horse", command.NewMacroCommand([]command.Command{
		walkUp, walkUp, walkRight,
	}))

	m.AddItem("status", "Prints Robot status (turned on/off, walk direction)", func() { printRobotInfo(r) })
	m.AddItem("stop", "Stops the Robot", r.Stop)

	m.ShowInstructions()
	m.Run()
}

func bindWalk(r *robot.Robot, dir int64) func() {
	return func() {
		r.Walk(dir)
	}
}

func directionToString(dir int64) string {
	switch dir {
	case robot.LEFT:
		return "left"
	case robot.RIGHT:
		return "right"
	case robot.UP:
		return "up"
	case robot.DOWN:
		return "down"
	default:
		return "stopped"
	}
}

func printRobotInfo(r *robot.Robot) {
	defer r.Output().Flush()
	if r.IsOn() {
		fmt.Fprintln(r.Output(), "Turned on")
		fmt.Fprintln(r.Output(), "Walk direction: ", directionToString(r.Direction()))

	} else {
		fmt.Fprintln(r.Output(), "Turned off")
	}
}
