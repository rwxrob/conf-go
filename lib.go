package conf

import (
	"io/fs"
	"os"
	"path"
)

// WriteMask is the value passed as the third argument to all file
// writes by default. The Map returned from NewMap uses this. (Other
// external implementations may not, however.)
var WritePerms fs.FileMode = 0600
var DirPerms fs.FileMode = 0700

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

// ExeDir returns the absolute path to a subdirectory of the home
// config directory named after the current executable binary.
// Will return empty string if either home directory or current
// executable name cannot be determined.
func ExeDir() string {
	abspath, _ := os.Executable()
	x := path.Base(abspath)
	dir := Home()
	if x == "" || dir == "" {
		return ""
	}
	return path.Join(dir, x)
}

// ExeDirFile returns an absolute path to the specified file within the
// Exe configuration directory. See Exe.
func ExeDirFile(file string) string {
	return path.Join(ExeDir(), file)
}
