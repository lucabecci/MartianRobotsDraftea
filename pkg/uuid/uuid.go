package uuid

import (
	"crypto/rand"
	"errors"
	"fmt"
)

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
