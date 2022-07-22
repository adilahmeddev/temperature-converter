package main

import (
	"bytes"
	"excercise4"
	"excercise4/assert"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

type Driver struct {
	path string
}

func (d Driver) ConvertToF(celsius float64) (float64, error) {
	out := &bytes.Buffer{}

	command := exec.Command(d.path)
	stdinPipe, err := command.StdinPipe()
	if err != nil {
		return 0, err
	}
	command.Stdout = out
	err = command.Start()

	_, err = fmt.Fprint(stdinPipe, "c\n")
	if err != nil {
		return 0, err
	}

	_, err = fmt.Fprintf(stdinPipe, "%.2f\n", celsius)
	if err != nil {
		return 0, err
	}

	err = command.Wait()

	lines := strings.Split(out.String(), "\n")
	fah := lines[len(lines)-1]

	if err != nil {
		return 0, err
	}
	float, err := strconv.ParseFloat(fah, 64)
	if err != nil {
		return 0, err
	}

	return float, nil
}

func (d Driver) ConvertToC(fah float64) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func TestName(t *testing.T) {
	cleanup, path, err := BuildBinary()
	assert.NoError(t, err)

	t.Cleanup(cleanup)

	excercise4.ConverterSpecification(t, Driver{
		path: path,
	})
}

const (
	baseBinName = "temp-converter-testbinary"
)

func BuildBinary() (cleanup func(), cmdPath string, err error) {
	binName := baseBinName

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		return nil, "", fmt.Errorf("cannot build tool %s: %s", binName, err)
	}

	build.Wait()

	dir, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	cmdPath = filepath.Join(dir, binName)

	cleanup = func() {
		os.Remove(binName)
	}

	return
}
