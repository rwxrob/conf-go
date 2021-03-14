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

// ConfigDir first determines the name of the current executable.
// Then it checks for the existence of any of the following (in order)
// and returns if found:
//
//    $XDG_CONFIG_HOME/<name>/
//    $HOME/.config/<name>/
//    $HOME/.<name>/
//
// If none are found ~/.config/<name>/ will be returned.
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
	_, err = os.Stat(confdir)
	if err == nil {
		return filepath.Join(confdir, name)
	}
	dir := filepath.Join(usr.HomeDir, "."+name)
	_, err = os.Stat(dir)
	if err == nil {
		return dir
	}
	dir = filepath.Join(usr.HomeDir, ".config", name)
	return dir
}
