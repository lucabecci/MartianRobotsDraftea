package calculation

import (
	"fmt"
	"strings"

	"github.com/lucabecci/MartianRobotsDraftea/internal/grid"
	"github.com/lucabecci/MartianRobotsDraftea/internal/robot"
)

func Process(instrucs []string) ([]robot.Robot, error) {
	var result []robot.Robot
	coordinates := strings.Split(instrucs[0], " ")
	grid, err := grid.GetGrid(coordinates)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = grid.Validate()
	if err != nil {
		fmt.Println(err.Error())
		return []robot.Robot{}, err
	}
	instrucs = append(instrucs[:0], instrucs[0+1:]...)
	for i := 0; i < len(instrucs); i++ {
		if (i % 2) == 0 {
			values := strings.Split(instrucs[i], " ")
			robot_instructs := strings.Split(instrucs[i+1], "")
			robot, err := robot.GetRobot(values[0], values[1], values[2])
			if err != nil {
				fmt.Println(err.Error())
			}
			robot.ExecInstruction(robot_instructs, grid)
			result = append(result, robot)
		}
	}
	return result, nil
}
