package file

import (
	"errors"
	"io/ioutil"
	"strings"
)

func ProcessFile(path string) ([]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("error trying to read the input file")
	}
	instrucs := strings.Split(string(bytes), "\n")
	return instrucs, nil
}
