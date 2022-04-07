package robot

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/lucabecci/MartianRobotsDraftea/internal/cardenal"
	"github.com/lucabecci/MartianRobotsDraftea/internal/grid"
	"github.com/lucabecci/MartianRobotsDraftea/internal/storage"
	"github.com/lucabecci/MartianRobotsDraftea/pkg/uuid"
)

type Robot struct {
	UUID        string
	X           int
	Y           int
	Orientation string
	Lost        bool
}

func GetRobot(x, y string, orientation string) (Robot, error) {
	coordinate_x, err := strconv.Atoi(x)
	if err != nil {
		return Robot{}, errors.New("error to create the robot because the coordinate X is invalid")
	}
	coordinate_y, err := strconv.Atoi(y)
	if err != nil {
		return Robot{}, errors.New("error to create the robot because the coordinate Y is invalid")
	}
	uuid, err := uuid.Generate()
	if err != nil {
		return Robot{}, err
	}
	robot := Robot{UUID: uuid, X: coordinate_x, Y: coordinate_y, Orientation: orientation, Lost: false}
	return robot, nil
}

func (r *Robot) ExecInstruction(instruction []string, grid grid.Grid) {
	for i := 0; i < len(instruction); i++ {
		if r.Lost {
			break
		}
		switch instruction[i] {
		case "F":
			r.moveForward(grid)
		case "R":
			r.moveRight()
		case "L":
			r.moveLeft()
		}
	}
}

func (r *Robot) moveLeft() {
	if cardenal.FindIndex(r.Orientation)-1 < 0 {
		r.Orientation = cardenal.Cardenal[len(cardenal.Cardenal)-1]
	} else {
		r.Orientation = cardenal.Cardenal[cardenal.FindIndex(r.Orientation)-1]
	}
}

func (r *Robot) moveForward(grid grid.Grid) {
	if r.Orientation == "N" {
		if r.Y == grid.Y {
			if !r.lostRobot() {
				r.Lost = true
				converted := r.convert()
				storage.Insert(converted)
			}
		} else {
			r.Y++
		}
	} else if r.Orientation == "S" {
		if r.Y == 0 {
			r.Lost = true
		} else {
			r.Y--
		}
	} else if r.Orientation == "E" {
		if r.X == grid.X {
			if !r.lostRobot() {
				r.Lost = true
				converted := r.convert()
				storage.Insert(converted)
				return
			}
		} else {
			r.X++
		}
	} else {
		if r.X == 0 {
			r.Lost = true
		} else {
			r.X--
		}
	}
}

func (r *Robot) moveRight() {
	if cardenal.FindIndex(r.Orientation)+1 >= len(cardenal.Cardenal) {
		r.Orientation = cardenal.Cardenal[0]
	} else {
		r.Orientation = cardenal.Cardenal[cardenal.FindIndex(r.Orientation)+1]
	}
}

func (r *Robot) lostRobot() bool {
	_, ubicationExists := storage.GetByAxis(r.X, r.Y)
	if ubicationExists {
		return true
	} else {
		return false
	}
}

func (r *Robot) convert() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(r)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}
