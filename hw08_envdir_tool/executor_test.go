package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env, err := ReadDir("./testdata/env")
	require.NoError(t, err)
	code := RunCmd([]string{"ls", "-l", "-a"}, env)

	require.Equal(t, code, 0)
}

func TestRunCmdWithError(t *testing.T) {
	code := RunCmd([]string{"mkdir", "./testdata"}, nil)

	require.Equal(t, code, 1)
}

func TestWithError(t *testing.T) {
	code := RunCmd([]string{"qwe", ""}, nil)

	require.Equal(t, code, -1)
}

func TestEnvVariablesError(t *testing.T) {
	env, err := ReadDir("./testdata/env")
	require.NoError(t, err)
	code := RunCmd([]string{"ls", "-la"}, env)

	require.Equal(t, code, 0)
	require.Equal(t, os.Getenv("BAR"), "bar")
	require.Equal(t, os.Getenv("EMPTY"), "")
	require.Equal(t, os.Getenv("FOO"), "   foo\nwith new line")
	require.Equal(t, os.Getenv("HELLO"), "\"hello\"")
	require.Equal(t, os.Getenv("UNSET"), "")
}

func TestReplaceEnvError(t *testing.T) {
	os.Setenv("UNSET", "not empty")
	os.Setenv("BAR", "not bar")

	require.Equal(t, os.Getenv("UNSET"), "not empty")
	require.Equal(t, os.Getenv("BAR"), "not bar")

	env, err := ReadDir("./testdata/env")
	require.NoError(t, err)
	code := RunCmd([]string{"ls", "-la"}, env)

	require.Equal(t, code, 0)
	require.Equal(t, os.Getenv("UNSET"), "")
	require.Equal(t, os.Getenv("BAR"), "bar")
}
