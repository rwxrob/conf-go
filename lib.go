package conf

import (
	"os"
	"path"
)

// Home returns the value of CONFIG_DIR environment variable if set, or
// the os.UserConfigDir.
//
func Home() string {
	dir := os.Getenv("CONFIG_DIR")
	if dir != "" {
		return dir
	}
	dir, _ = os.UserConfigDir()
	return dir
}

// Exe returns the absolute path to a subdirectory of the home
// config directory named after the current executable binary.
// Will return empty string if either home directory or current
// executable name cannot be determined.
//
func Exe() string {
	abspath, _ := os.Executable()
	x := path.Base(abspath)
	dir := Home()
	if x == "" || dir == "" {
		return ""
	}
	return path.Join(dir, x)
}

// ExeFile returns an absolute path to the specified file within the Exe
// directory.
func ExeFile(file string) string {
	return path.Join(Exe(), file)
}
