package conf

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
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
// WARNING: use of line returns will *not* be escaped automatically and
// will break the persistent file that contains the configuration
// data unless you escape them yourself.
type Setter interface {
	Set(key string, val string) error
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

// Mutex implements a mutex with Lock and Unlock that should lock all
// reads and writes while locked. Implementations may decide to limit
// scope to same runtime memory space, or to extend to a system-level
// scope that other external running processes can also observe.
type Mutex interface {
	Lock()
	Unlock()
}

// Map implements a conf.Map suitable for returning from NewMap. See the
// individual (single-method) interface descriptions for details.
//
// Stringer must be implemented as a Parse-able string (see Parse).
//
// Raw returns the inner map[string]string for direct manipulation that
// bypasses locking such as when ranging or sorting.
type Map interface {
	Mutex
	Getter
	Setter
	JSONify
	fmt.Stringer
	Raw() map[string]string
}

type mapStruct struct {
	sync.Mutex
	m map[string]string
}

// NewMap returns a new struct that fulfills the Map interface and
// embeds a sync.Mutex to fulfill the conf.Mutex interface. Eventually,
// the Mutex implementation may be expanded to allow other processes to
// observe the lock as well.
func NewMap() *mapStruct {
	m := new(mapStruct)
	m.m = map[string]string{}
	return m
}

// Parse constructs and returns a struct that fulfills the Map interface
// from parsed bytes. The data being parsed must comply with the
// requires for configuration data:
//
// * One configuration key=value pair per line
// * Each line must contain an equal sign (=) delimiter
// * Whitespace around equal delim is NOT ignored
// * Lines must end with standard ending (\r?\n)
// * Keys and values may be any string value (except line ending)
// * Editable string keys and values are strongly recommended
// * Empty lines will be ignored and overwritten (by Write, etc.)
// * No support for comments
//
func Parse(b []byte) (*mapStruct, error) {
	m := NewMap()
	scanner := bufio.NewScanner(bytes.NewReader(b))
	n := 1
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, "=")
		if len(f) != 2 {
			return nil, fmt.Errorf("invalid config (line %v): %v\n", n, line)
		}
		m.m[f[0]] = f[1]
		n++
	}
	return m, nil
}

// Read returns a new struct that fulfills Map interface by reading the
// default location (see ExeDirFile passed "values").
func Read() (*mapStruct, error) {
	buf, err := os.ReadFile(ExeDirFile("values"))
	if err != nil {
		return nil, err
	}
	return Parse(buf)
}

// Write locks the Map and it to the default location (ExeDirFiles passed
// "values" argument) with the default package permissions (see
// WritePerms).
func Write(m Map) error {
	m.Lock()
	defer m.Unlock()
	err := os.MkdirAll(ExeDir(), DirPerms)
	if err != nil {
		return err
	}
	return os.WriteFile(ExeDirFile("values"), []byte(m.String()), WritePerms)
}

func (m *mapStruct) Raw() map[string]string { return m.m }

func (m *mapStruct) Get(key string) string {
	m.Lock()
	defer m.Unlock()
	return m.m[key]
}

func (m *mapStruct) Set(key, val string) error {
	m.Lock()
	m.m[key] = val
	m.Unlock()
	return Write(m)
}

func (m *mapStruct) Keys() []string {
	keys := []string{}
	for k, _ := range m.m {
		keys = append(keys, k)
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
