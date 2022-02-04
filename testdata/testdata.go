package testdata

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

func Load(filepath string, s interface{}) error {
	dir := ""
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir = path.Join(path.Dir(filename))
	}
	fullPath := path.Join(dir, filepath)
	body, err := os.ReadFile(fullPath)
	if err != nil {
		log.Println("Error helper.LoadFile: unable to read file", err)
		panic(err)
	}

	buffer := new(bytes.Buffer)
	_, er := buffer.Write(body)
	if er != nil {
		return er
	}

	if buffer.Len() == 0 {
		return errors.New("testdata is empty")
	}

	if er := json.NewDecoder(buffer).Decode(s); er != nil {
		return er
	}

	return nil
}
