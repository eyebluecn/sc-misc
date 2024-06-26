package util

import (
	"runtime"
)

// whether windows develop environment
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
func IsMac() bool {
	return runtime.GOOS == "darwin"
}
