package uuid_test

import (
	"testing"

	"github.com/lucabecci/MartianRobotsDraftea/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("Generate 5 uuid", func(t *testing.T) {
		for i := 0; i < 6; i++ {
			_, err := uuid.Generate()
			assert.Equal(t, err, nil)
		}
	})

	t.Run("Generate 25 uuid", func(t *testing.T) {
		for i := 0; i < 26; i++ {
			_, err := uuid.Generate()
			assert.Equal(t, err, nil)
		}
	})
}
