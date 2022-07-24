package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ErrUnsupportedFile = errors.New("unsupported file")
	ErrOpeningFile     = errors.New("offset exceeds file size")
	ErrReadingDir      = errors.New("directory is unreadable")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, ErrReadingDir
	}

	environment := make(Environment)
	for _, file := range files {
		var env EnvValue
		env.Value = ""
		env.NeedRemove = false

		fileName := dir + "/" + file.Name()

		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			return nil, ErrOpeningFile
		}

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println(err)
			return nil, ErrUnsupportedFile
		}
		fileSize := fileInfo.Size()
		if fileSize == 0 {
			env.NeedRemove = true
		}

		r := bufio.NewReader(f)
		s, _ := Readln(r)
		if s == "" {
			env.NeedRemove = true
		}

		env.Value = s

		for _, envVar := range os.Environ() {
			pair := strings.SplitN(envVar, "=", 2)
			if pair[0] == env.Value {
				os.Unsetenv(pair[0])
			}
		}
		environment[file.Name()] = env
	}

	return environment, nil
}

func Readln(r *bufio.Reader) (string, error) {
	isPrefix := true
	var err error
	var line []byte
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		line = bytes.ReplaceAll(line, []byte{0}, []byte{'\n'})
		line = []byte(strings.TrimRight(string(line), " "))
	}
	return string(line), err
}
