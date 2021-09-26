package conf

import (
	"io/fs"
	"strings"
)

// WriteMask is the value passed as the third argument to all file
// writes by default. The Map returned from NewMap uses this. (Other
// external implementations may not, however.)
var WritePerms fs.FileMode = 0600
var DirPerms fs.FileMode = 0700

// Escape converts actual carriage returns and line returns into their
// respective \r and \n equivalents suitable as values (or keys,
// although discouraged) passed to Map.Set.
func Escape(s string) string {
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	return s
}
