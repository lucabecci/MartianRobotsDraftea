package file_test

import (
	"testing"

	"github.com/lucabecci/MartianRobotsDraftea/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestProcessFile(t *testing.T) {
	t.Run("Generate instructions based on a file", func(t *testing.T) {
		_, err := file.ProcessFile("../../input.txt")
		assert.Equal(t, err, nil)
	})
}
