package cardenal_test

import (
	"testing"

	"github.com/lucabecci/MartianRobotsDraftea/internal/cardenal"
	"github.com/stretchr/testify/assert"
)

func TestGetGrid(t *testing.T) {
	t.Run("Get cardenal N", func(t *testing.T) {
		coordinate := "N"
		index := cardenal.FindIndex(coordinate)
		assert.Equal(t, index, 0)
	})
	t.Run("Get cardenal E", func(t *testing.T) {
		coordinate := "E"
		index := cardenal.FindIndex(coordinate)
		assert.Equal(t, index, 1)
	})
	t.Run("Get cardenal S", func(t *testing.T) {
		coordinate := "S"
		index := cardenal.FindIndex(coordinate)
		assert.Equal(t, index, 2)
	})
	t.Run("Get cardenal W", func(t *testing.T) {
		coordinate := "W"
		index := cardenal.FindIndex(coordinate)
		assert.Equal(t, index, 3)
	})
}
