package calculation_test

import (
	"testing"

	"github.com/lucabecci/MartianRobotsDraftea/internal/calculation"
	"github.com/lucabecci/MartianRobotsDraftea/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("Process with the input file", func(t *testing.T) {
		input, _ := file.ProcessFile("../../input.txt")
		result, _ := calculation.Process(input)
		assert.Equal(t, result[0].Lost, false)
		assert.Equal(t, result[1].Lost, true)
		assert.Equal(t, result[2].Lost, false)
	})
}
