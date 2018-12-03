package util

import (
	"github.com/liamylian/jsontime"
	"io"
	"io/ioutil"
)

var cusJson = jsontime.ConfigWithCustomTimeFormat

func Unmarshal(r io.ReadCloser, v interface{}) error {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = cusJson.Unmarshal(bytes, v)
	return err
}

func Marshal(v interface{}) ([]byte, error) {
	return cusJson.Marshal(v)
}
