package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"sync"
)

// Getter is implemented by anything with a string value to return.
type Getter interface {
	Get(key string) string
}

// Setter is implemented by anything that takes a string value with
// a string key. Any valid string data is allowed as the key (including
// periods, dashes, brackets, even emojis).  Set sets the value of key
// to the string.
//
// Warning: use of line returns will *not* be escaped automatically and
// will break other methods that depend on scan.ScanLine line endings.
// This is to preserve the performance of a single character delimiter
// per line without any additional overhead.
type Setter interface {
	Set(key string, val string) error
}

// Reader is implemented by anything that takes an io.Reader and reads
// from it return any error in the process. The data being read must
// comply with the requires for configuration data:
//
// * One configuration key=value pair per line
// * Each line must contain an equal sign (=) delimiter
// * Whitespace around equal delim is NOT ignored
// * Lines must end with scanner.ScanLine endings (\r?\n)
// * Keys and value may be any string value (except line ending)
// * Editable string keys and values are strongly recommended
// * Empty lines will be ignored and overwritten (by Save, etc.)
// * No support for comments
//
type Reader interface {
	Read(r io.Reader) error
}

// Saver is implemented by anything that saves to a default or inferred
// location (which must be the value returned by calling the ExeFile
// function and passing it the "values" argument. See ExeFile for
// details. Calling ExeFile specifically, however, is not required by
// this interface.)
type Saver interface {
	Save() error
}

// Loader is implemented by anything that loads from a default or
// inferred location. See Saver interface for details.
type Loader interface {
	Load() error
}

// Editable is implemented by opening a configured or default editor set
// by the user either usually as a configuration or environment variable
// (such as EDITOR or VISUAL). Implementations will differ depending on
// operating system supported. Implementations may decide to hand off
// process control to the editor program in a way that does not return
// control to the calling program (such as with UNIX/Linux exec). Please
// make not of such in any implementation documentation.
type Editable interface {
	Edit() error
}

// JSONify implements a JSON method to return a string of valid JSON
// (compressed or indented). If an error occurs it must be returned as
// the ERROR key with the message as value in valid JSON.
//
//     {"ERROR": "some error message"}
//
// Please document whether the JSON output will be compressed on
// a single line or not.
type JSONify interface {
	JSON() string
}

// Map implements a conf.Map suitable for returning from NewMap. See the
// individual (single-method) interface descriptions for details.
//
// Stringer must be implemented as a Load/Save-able string.
//
// Raw returns the inner map[string]string for direct manipulation that
// bypasses locking such as when ranging or sorting.
type Map interface {
	Getter
	Setter
	Saver
	Loader
	Editable
	JSONify
	fmt.Stringer
	Raw() map[string]string
}

type mapStruct struct {
	sync.RWMutex
	m map[string]string
}

// NewMap returns a new struct that fulfills the Map interface and
// embeds a sync.RWMutex.
func NewMap() *mapStruct {
	m := new(mapStruct)
	m.m = map[string]string{}
	return m
}

func (m *mapStruct) Raw() map[string]string { return m.m }

func (m *mapStruct) Get(key string) string {
	m.RLock()
	defer m.RUnlock()
	return m.m[key]
}

func (m *mapStruct) Set(key, val string) error {
	m.RLock()
	defer m.RUnlock()
	m.m[key] = val
	// TODO add line return escapes
	return m.Save()
}

func (m *mapStruct) Save() error {
	// TODO
	return nil
}

func (m *mapStruct) Read(src io.Reader) error {
	// TODO
	return nil
}

func (m *mapStruct) Load() error {
	// TODO read the implied ExeFile values
	return nil
}

func (m *mapStruct) Edit() error {
	// TODO
	return nil
}

func (m *mapStruct) Keys() []string {
	keys := make([]string, len(m.m))
	n := 0
	for k, _ := range m.m {
		keys[n] = k
		n++
	}
	sort.Strings(keys)
	return keys
}

func (m mapStruct) String() string {
	buf := ""
	for _, k := range m.Keys() {
		buf = buf + fmt.Sprintf("%v=%v\n", k, m.m[k])
	}
	return buf
}

func (m mapStruct) JSON() string {
	byt, err := json.Marshal(m.m)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":%q}", err)
	}
	return string(byt)
}

func (m mapStruct) Print()     { fmt.Print(m) }
func (m mapStruct) PrintJSON() { fmt.Println(m.JSON()) }
