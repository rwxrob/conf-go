package conf

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

// ErrorNewer is an error that prints "Newer config file detected" but
// can be changed to any other error including those written in other
// languages.
var ErrorNewer = errors.New("Newer config file detected.")

// ConfigDir looks for $XDG_CONFIG_HOME and if set looks within it to
// resolve the path to the named configuration directory. If that
// environment variable is not set will also look for the ~/.config
// directory and do the same. If neither of these are present will look
// for a directory beginning with a dot in the home directory itself. If
// none of these are present, then a ~/.config directory will be created
// within the home directory and a named directory created within it.
// The name added to the directory path (once resolved) is the base name
// of the current running executable (os.Executable()). Returns empty
// string if any error is encountered.
func ConfigDir() string {
	name, err := os.Executable()
	if err != nil {
		return ""
	}
	name = filepath.Base(name)
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	confdir := os.Getenv("XDG_CONFIG_HOME")
	if confdir != "" {
		return filepath.Join(confdir, name)
	}
	confdir = filepath.Join(usr.HomeDir, ".config")
	if exists(confdir) {
		return filepath.Join(confdir, name)
	}
	dir := filepath.Join(usr.HomeDir, "."+name)
	if exists(dir) {
		return dir
	}
	dir = filepath.Join(usr.HomeDir, ".config", name)
	return dir
}
