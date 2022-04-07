package file

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func ProcessFile() ([]string, error) {
	path, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(path + "/input.txt")
	if err != nil {
		return nil, errors.New("error trying to read the input file")
	}
	instrucs := strings.Split(string(bytes), "\n")
	return instrucs, nil
}
