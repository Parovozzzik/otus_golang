package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var envDir = "testdata/env"

func TestReadDir(t *testing.T) {
	envMap, err := ReadDir(envDir)
	require.NoError(t, err)
	require.Equal(t, EnvValue{Value: "bar", NeedRemove: false}, envMap["BAR"])
	require.Equal(t, EnvValue{Value: "", NeedRemove: true}, envMap["EMPTY"])
	require.Equal(t, EnvValue{Value: "   foo\nwith new line", NeedRemove: false}, envMap["FOO"])
	require.Equal(t, EnvValue{Value: "\"hello\"", NeedRemove: false}, envMap["HELLO"])
	require.Equal(t, EnvValue{Value: "", NeedRemove: true}, envMap["UNSET"])
}

func TestEmptyDir(t *testing.T) {
	envMap, err := ReadDir("testdata/no_dir")
	require.Error(t, err, "open testdata/nodir: no such file or directory")
	require.Equal(t, 0, len(envMap))
}

func TestEqualSymbol(t *testing.T) {
	file, err := os.CreateTemp(envDir, "QWE_*")
	require.NoError(t, err)
	file.Write([]byte("test"))
	file.Close()
	defer os.Remove(file.Name())
	envMap, err := ReadDir(envDir)
	require.NoError(t, err)
	require.Equal(t, EnvValue{Value: "test", NeedRemove: false}, envMap[strings.ReplaceAll(file.Name(), envDir+"/", "")])
}
