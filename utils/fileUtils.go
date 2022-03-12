package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	if IsEmpty(fileName) {
		return "", errors.New("boş veri dosya adı olarak kullanılamaz")
	}

	bytes, err := ioutil.ReadFile(fileName)
	checkErr(err)

	return string(bytes), nil
}
