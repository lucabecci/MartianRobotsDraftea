package main

import (
	"fmt"
	"os"

	"github.com/lucabecci/MartianRobotsDraftea/internal/calculation"
	"github.com/lucabecci/MartianRobotsDraftea/pkg/file"
)

func main() {
	data, err := file.ProcessFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	robots, err := calculation.Process(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	for x, robot := range robots {
		fmt.Println("ROBOT", x, robot)
	}
}
