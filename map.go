package conf

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"sync"
)

// Getter is implemented by anything with a string value to return. If
// the key does not exist an empty string must be returned.
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

// HasHome implements a meta variable Home to contain the full path to
// the configuration home directory containing a subdirectory named
// exactly to match HasName. By default this value must be the returned
// value of os.UserConfigDir.
type HasHome interface {
	Home() string
	SetHome(s string)
}

// HasName implements a meta variable Name to uniquely identify one Map
// from another. By default this value must be the base name of the
// current os.Executable.
type HasName interface {
	Name() string
	SetName(s string)
}

// HasFile implements a meta variable File to contains the name of the
// file within HasName subdirectory containing the key=value data.
// The default must be 'values'.
type HasFile interface {
	File() string
	SetFile(s string)
}

// Reader implements a Read method that reads and initializes the
// internal struct from its default source (Map.Path in this case.)
type Reader interface {
	Read() error
}

// Writer implements a Write method that writes the internal struct
// state to a default file location (Map.Path in this case).
type Writer interface {
	Write() error
}

// Parser implements a Parse method that parses data and adds the
// key=value pairs to the internal struct from parsed bytes.
// The Parse method must be additive and must overwrite any existing
// values already set. It must not delete any existing values that are
// not being overwritten. The data being parsed must comply with the
// following requirements for configuration data:
//
//     * One configuration key=value pair per line
//     * Each line must contain an equal sign (=) delimiter
//     * Whitespace around equal delim is NOT ignored
//     * Lines must end with standard ending (\r?\n)
//     * Keys and values may be any string value (except line ending)
//     * Editable string keys and values are strongly recommended
//     * Empty lines will be ignored and overwritten (by Write, etc.)
//     * No support for comments
//
type Parser interface {
	Parse(b []byte) error
}

// Deleter implements a Delete method that removes a given key=value
// pair from the internal struct and persists it. Deleter is a no-op,
// meaning it doesn't return anything. It either deletes it or not.
type Deleter interface {
	Delete(key string)
}

// Editor implements an Edit method that will detect the best command
// line editor available and pass the full path to the configuration
// file (see Path). All implementations must first look within for an
// EDITOR key, followed by the EDITOR and VISUAL environment variables.
//
//     Map.Get("EDITOR")
//     os.Getenv("EDITOR")
//     os.Getenv("VISUAL")
//
type Editor interface {
	Edit() error
}

// Map implements a conf.Map suitable for returning from NewMap. See the
// individual (single-method) interface descriptions for details.
//
// Stringer must be implemented as a Parse-able string (see Parse).
//
// Raw returns the inner map[string]string for direct manipulation that
// bypasses locking such as when ranging or sorting.
type Map interface {
	HasHome
	HasName
	HasFile
	Mutex
	Getter
	Setter
	Deleter
	Reader
	Writer
	Parser
	Editor
	JSONify
	fmt.Stringer
	Raw() map[string]string
}

// ----------------------------- mapStruct ----------------------------

type mapStruct struct {
	sync.Mutex
	home string
	name string
	file string
	m    map[string]string
}

// NewMap returns a new struct that fulfills the Map interface.
//
// The Edit method implemented will first look for an EDITOR key within
// its internal map and then look for the EDITOR and VISUAL environment
// variables (as is traditional on UNIX-based systems). The lookup of
// the map key allows configurations to persist the preferred editor for
// that specific configuration and file.
//
// The returned struct embeds a sync.Mutex to fulfill the Mutex
// interface. Eventually, the Mutex implementation may be expanded to
// allow other external processes to observe the lock as well.
//
func NewMap() *mapStruct {
	var err error
	m := new(mapStruct)
	m.m = map[string]string{}
	// m.Home()
	m.home, err = os.UserConfigDir()
	if err != nil {
		log.Println(err)
	}
	// m.Name()
	exe, _ := os.Executable()
	m.name = path.Base(exe)
	// m.File()
	m.file = "values"
	return m
}

func (m *mapStruct) Home() string           { return m.home }
func (m *mapStruct) SetHome(s string)       { m.home = s }
func (m *mapStruct) Name() string           { return m.name }
func (m *mapStruct) SetName(s string)       { m.name = s }
func (m *mapStruct) File() string           { return m.file }
func (m *mapStruct) SetFile(s string)       { m.file = s }
func (m *mapStruct) Raw() map[string]string { return m.m }
func (m *mapStruct) Print()                 { fmt.Print(m) }
func (m *mapStruct) PrintJSON()             { fmt.Println(m.JSON()) }

func (m *mapStruct) Dir() string {
	return path.Join(m.home, m.name)
}

func (m *mapStruct) Path() string {
	return path.Join(m.home, m.name, m.file)
}

func (m *mapStruct) Get(key string) string {
	m.Read()
	m.Lock()
	defer m.Unlock()
	if v, has := m.m[key]; has {
		return v
	}
	return ""
}

func (m *mapStruct) Set(key, val string) error {
	m.Read()
	m.Lock()
	m.m[key] = val
	m.Unlock()
	return m.Write()
}

func (m *mapStruct) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, key)
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

func (m *mapStruct) Read() error {
	buf, err := os.ReadFile(m.Path())
	if err != nil {
		return err
	}
	return m.Parse(buf)
}

func (m *mapStruct) Write() error {
	m.Lock()
	defer m.Unlock()
	err := os.MkdirAll(path.Join(m.home, m.name), DirPerms)
	if err != nil {
		return err
	}
	return os.WriteFile(m.Path(), []byte(m.String()), WritePerms)
}

func (m *mapStruct) Parse(b []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	n := 1
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, "=")
		if len(f) != 2 {
			return fmt.Errorf("invalid config (line %v): %v\n", n, line)
		}
		m.m[f[0]] = f[1]
		n++
	}
	return nil
}

func (m *mapStruct) getEditor() string {
	e := m.Get("EDITOR")
	if e != "" {
		return e
	}
	e = os.Getenv("EDITOR")
	if e != "" {
		return e
	}
	e = os.Getenv("VISUAL")
	if e != "" {
		return e
	}
	path, err := exec.LookPath("vi")
	if err != nil {
		return path
	}
	return ""
}

func (m *mapStruct) Edit() error {
	editor := m.getEditor()
	if editor == "" {
		return fmt.Errorf("unable to determine editor")
	}
	path, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	cmd := exec.Command(path, m.Path())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
