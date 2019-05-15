// Package osutil provides some cross-platform utilities for accessing
// system-dependent functions.
package osutil

// LogDir returns the canonical place to place log files for an app.
func LogDir(appname string) string {
	return logdir(appname)
}

// ConfigDir returns the canonical place to place configuration files for an
// app.
func ConfigDir(appname string) string {
	return configdir(appname)
}

// DataDir returns the canonical place to place application data for an app.
func DataDir(appname string) string {
	return datadir(appname)
}

// Open opens the given string using the system's default program delegation
// method.
func Open(object string) error {
	return open(object)
}

// GetSerialPorts returns a list of available serial ports on the system.
func GetSerialPorts() ([]string, error) {
	return getSerialPorts()
}
