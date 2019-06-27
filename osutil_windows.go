package osutil

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	shell32        = windows.NewLazyDLL("shell32.dll")
	shellExecute   = shell32.NewProc("ShellExecuteW")
	kernel32       = syscall.MustLoadDLL("kernel32.dll")
	queryDosDevice = kernel32.MustFindProc("QueryDosDeviceA")
)

func logdir(appname string) string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), appname, "logs")
}

func configdir(appname string) string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), appname)
}

func datadir(appname string) string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), appname, "data")
}

func open(object string) error {
	lpOperation, err := windows.UTF16PtrFromString("open")
	if err != nil {
		return err
	}

	lpFile, err := windows.UTF16PtrFromString(object)
	if err != nil {
		return err
	}

	_, _, err = shellExecute.Call(
		0,
		uintptr(unsafe.Pointer(lpOperation)),
		uintptr(unsafe.Pointer(lpFile)),
		0,
		0,
		0,
	)

	return err
}

func getSerialPorts() ([]string, error) {
	var (
		buf = make([]byte, 4096)
		n   int
	)

	for {
		ret, _, err := queryDosDevice.Call(0, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
		if ret == 0 {
			errno := err.(syscall.Errno)
			if errno == windows.ERROR_INSUFFICIENT_BUFFER {
				buf = make([]byte, len(buf)+4096)
				continue
			}
			return nil, err
		} else {
			n = int(ret)
			break
		}
	}

	devs := strings.Split(string(buf[:n-1]), "\x00")

	if len(devs) == 0 {
		return nil, errors.New("empty device list")
	}

	coms := make([]string, 0, len(devs))
	for _, dev := range devs {
		if strings.HasPrefix(dev, "COM") {
			coms = append(coms, dev)
		}
	}

	return coms, nil
}
