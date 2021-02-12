/*
Package conf is a simple package to assist with writing flat
configuration and other data to JSON files at an expected location
safely being sure not overwrite any configuration file that has been
updated since it was last written by the running program. This helps
avoid contention between two runtimes that might try to write to the
same file.

The JSON data is always saved as a one-dimensional, flat map of strings
with concurrency-safe Get() and Set() methods. Depth may be achieved by
using dotted or dashed notation for the keys but ultimately everything
is saved in a single map of key/value pairs where the keys and values
are both quoted UTF-8 strings. This improves efficiency and avoids type
inference issues.

*/
package conf
