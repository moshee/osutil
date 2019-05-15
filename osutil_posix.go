//+build !windows

package osutil

import (
	"errors"
	"os/exec"
	"path/filepath"
)

func getSerialPorts() ([]string, error) {
	mm, err := filepath.Glob(posixSerialGlobPattern)
	if err != nil {
		return nil, err
	}

	if len(mm) == 0 {
		return nil, errors.New("no serial devices found in /dev")
	}

	return mm, nil
}

func open(object string) error {
	cmd := exec.Command(openCommand, object)
	return cmd.Run()
}
