package grid

import (
	"errors"
	"strconv"
)

// Struct for a Grid
type Grid struct {
	X int
	Y int
}

// Create a grid from the coordinates
// Returns the grid created together with a possible error
func GetGrid(coordinates []string) (Grid, error) {
	coordinate_x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		return Grid{}, errors.New("the coordinate X is invalid")
	}
	coordinate_y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		return Grid{}, errors.New("the coordinate X is invalid")
	}

	return Grid{
		X: coordinate_x,
		Y: coordinate_y,
	}, nil
}

// Method of Grid struct
// Validates that the grid does not have an invalid x-axis and y-axis
// Returns an error if it exists
func (g *Grid) Validate() error {
	if g.X > 50 || g.Y > 50 {
		return errors.New("the grid is invalid")
	} else {
		return nil
	}
}
