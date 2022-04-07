package grid_test

import (
	"testing"

	"github.com/lucabecci/MartianRobotsDraftea/internal/grid"
	"github.com/stretchr/testify/assert"
)

func TestGetGrid(t *testing.T) {
	t.Run("Create a grid - without errors 1", func(t *testing.T) {
		test_case := [2]string{"2", "2"}
		grid, _ := grid.GetGrid(test_case[:])
		assert.Equal(t, grid.X, 2)
		assert.Equal(t, grid.Y, 2)
	})

	t.Run("Create a grid - without errors 2", func(t *testing.T) {
		test_case := [2]string{"22", "30"}
		grid, _ := grid.GetGrid(test_case[:])
		assert.Equal(t, grid.X, 22)
		assert.Equal(t, grid.Y, 30)
	})
}

func TestValidate(t *testing.T) {
	t.Run("Validate a grid with errors", func(t *testing.T) {
		test_case := [2]string{"51", "61"}
		grid, _ := grid.GetGrid(test_case[:])
		assert.Equal(t, grid.X, 51)
		assert.Equal(t, grid.Y, 61)
		err := grid.Validate()
		assert.Equal(t, err.Error(), "the grid is invalid")
	})

	t.Run("Validate a grid without errors", func(t *testing.T) {
		test_case := [2]string{"2", "3"}
		grid, _ := grid.GetGrid(test_case[:])
		assert.Equal(t, grid.X, 2)
		assert.Equal(t, grid.Y, 3)
		err := grid.Validate()
		assert.Equal(t, err, nil)
	})
}
