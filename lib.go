package conf

import (
	"io/fs"
)

// WriteMask is the value passed as the third argument to all file
// writes by default. The Map returned from NewMap uses this. (Other
// external implementations may not, however.)
var WritePerms fs.FileMode = 0600
var DirPerms fs.FileMode = 0700
