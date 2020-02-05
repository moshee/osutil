package osutil

import (
	"os"
	"path/filepath"
)

const (
	posixSerialGlobPattern = "/dev/ttyACM*"
	openCommand            = "xdg-open"
)

func logdir(appname string) string {
	return filepath.Join(os.Getenv("XDG_CACHE_HOME"), appname)
}

func configdir(appname string) string {
	return filepath.Join(os.Getenv("XDG_CONFIG_HOME"), appname)
}

func datadir(appname string) string {
	return filepath.Join(os.Getenv("XDG_DATA_HOME"), appname)
}
