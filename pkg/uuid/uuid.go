package uuid

import (
	"crypto/rand"
	"errors"
	"fmt"
)

// Generate function that will allow us to generate a unique identifier
// It returns the unique identifier and an error if it exists
func Generate() (string, error) {
	var uuid string
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New("error to generate the uuid")
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}
