package conf

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrorNewer is an error that prints "Newer config file detected" but
// can be changed to any other error including those written in other
// languages.
var ErrorNewer = errors.New("Newer config file detected.")

// ConfigDir first determines the name of the current executable
// (os.Executable()).
//
// Then, the existence of any of the following are checked (in order)
// and returned if found:
//
//    $XDG_CONFIG_HOME/<name>
//    $HOME/.config/<name>/
//    $HOME/.<name>/
//
// If none are found $HOME/.config/<name>/ will be returned (whether or not
// it exists).
//
// Note: $HOME is simply a visual indicator of os.UserHomeDir() and will not
// always match the environment variable directly.
func ConfigDir() string {
	name, err := os.Executable()
	if err != nil {
		return ""
	}
	name = filepath.Base(name)
	confdir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	// checking if confdir exist because if $XDG_CONFIG_HOME does not exist
	// then confdir will be == $HOME/.config and this folder might not exist
	_, err = os.Stat(confdir)
	if err == nil {
		return filepath.Join(confdir, name)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	dir := filepath.Join(homedir, "."+name)
	_, err = os.Stat(dir)
	if err == nil {
		return dir
	}
	dir = filepath.Join(homedir, ".config", name)
	return dir
}
