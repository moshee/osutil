package osutil

import (
	"fmt"
	"os/user"
	"path/filepath"
)

const (
	posixSerialGlobPattern = "/dev/tty.usbmodem*"
	openCommand            = "open"
)

var u *user.User

func init() {
	var err error
	u, err = user.Current()
	if err != nil {
		panic(fmt.Errorf("can't get current user: %v", err))
	}
}

func logdir(appname string) string {
	return filepath.Join(u.HomeDir, "Library", "Logs", appname)
}

func configdir(appname string) string {
	return filepath.Join(u.HomeDir, "Library", "Preferences", appname)
}

func datadir(appname string) string {
	return filepath.Join(u.HomeDir, "Library", "Application Support", appname)
}
