package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// Config encapsulates the Data structure and provides additional meta
// data such as a timestamp when last Saved and the absolute path to the
// directory containing the configuration File. Config also includes
// a sync.RWMutex for performant concurrency safety.
type Config struct {
	mu      sync.RWMutex
	unsaved bool

	// Absolute path to the directory containing File. Defaults to the
	// output of ConfigDir.
	Dir string `json:"dir,omitempty"`

	// File name defaults to "config.json", which will be placed into the
	// Dir directory.
	File string `json:"file,omitempty"`

	// Data is the actual key/value data.
	Data Data `json:"data"`

	// Save is set to the last time a save successfully completed.
	Saved time.Time `json:"saved,omitempty"`

	// Updated is set every time Set is called.
	Updated time.Time `json:"updated,omitempty"`
}

// Path returns the absolute path to the configuration file. Returns
// empty Path if an error occurs. Does not check for existence of file.
func (jc *Config) Path() string {
	dir, err := filepath.Abs(jc.Dir)
	if err != nil {
		return ""
	}
	return filepath.Join(dir, jc.File)
}

// New accepts up to three variadic arguments and returns a new Config
// pointer. First argument is the Dir path string, which will be
// converted to an absolute path if passed a relative directory. Second
// is File name.  Note that New() does not automatically load content
// from Path() nor does it create a new configuration directory and
// file. Use Load() when needed instead. Panics if more than two
// arguments are passed.
func New(args ...string) *Config {
	var err error
	jc := new(Config)
	jc.Data = map[string]string{}

	switch len(args) {
	case 0:
		jc.Dir = ConfigDir()
		jc.File = "config.json"
	case 1:
		jc.Dir, err = filepath.Abs(args[0])
		jc.File = "config.json"
	case 2:
		jc.Dir, err = filepath.Abs(args[0])
		jc.File = args[1]
	default:
		err = errors.New("too many arguments")
	}

	if err != nil {
		panic(err)
	}
	return jc
}

// NewFromJSON will return a new Config pointer by setting its Data to
// that which is Unmarshaled from a JSON byte array. The rest of the
// arguments are the same as New().
func NewFromJSON(jsn []byte, args ...string) (*Config, error) {
	jc := New(args...)
	err := json.Unmarshal(jsn, &jc)
	return jc, err
}

// NewFromFile uses the specific path to read a json file. However, the
// path does not automatically set the Dir and File, which must be
// provided as separate args if the default initial values are
// not wanted. This is occasionally useful when a file is already
// available but not in the preferred confirugration location.
func NewFromFile(path string, args ...string) (*Config, error) {
	byt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewFromJSON(byt, args...)
}

// TODO:
// func Import
// func ImportFile
// func ImportJSON

// Get returns the value for given key in a concurrency-safe way.
func (jc *Config) Get(key string) string {
	jc.mu.RLock()
	defer jc.mu.RUnlock()
	return jc.Data[key]
}

// Set assigns the value to key but does not Save(), which must be done
// explicitly. Call SetSave() or SetForceSave() instead when an
// immediate Save() is required. Set is safe for concurrency.
func (jc *Config) Set(key string, val string) {
	jc.mu.Lock()
	defer jc.mu.Unlock()
	jc.Data[key] = val
	jc.Updated = time.Now()
}

// SetSave calls Set() and then Save() as a convienience.
func (jc *Config) SetSave(key string, val string) error {
	jc.Set(key, val)
	return jc.Save()
}

// SetForceSave calls Set() and then ForceSave() as a convienience.
func (jc *Config) SetForceSave(key string, val string) error {
	jc.Set(key, val)
	return jc.ForceSave()
}

// Init replaces the Data from the current object with a new
// map[string]string and creates new configuration file at Path()
// containing nothing but updated meta data (data is empty) returning an
// error if something goes wrong. WARNING: If a file already exists,
// Init() will purge the data from the file.
func (jc *Config) Init() error {
	jc.Data = map[string]string{}
	return jc.ForceSave()
}

// Load initializes the Config object with data freshly loaded from the
// current Path() throwing away the internal reference to any previous
// Data (which will be cleaned up with normal garbage collection). This
// is useful for situations where reloading from the last Save() is
// wanted. If no configuration file exists, calling Load will initialize
// a new one at the Path() location.
func (jc *Config) Load() error {

	if !exists(jc.Path()) {
		err := jc.Init()
		if err != nil {
			return err
		}
	}

	newjc, err := NewFromFile(jc.Path())
	if err != nil {
		return err
	}
	jc.Data = newjc.Data
	jc.Saved = newjc.Saved
	jc.Updated = newjc.Updated
	return nil
}

// Save checks the Saved time within file at Path() and compares it to
// Updated refusing to overwrite the file if Updated is older than the
// last save (which would create an over-written changes situation).
// ErrorNewer is returned in such cases.  Creates any parent directories
// as needed along with a new Path() file if one does not already exist.
func (jc *Config) Save() error {
	jc.mu.RLock()
	defer jc.mu.RUnlock()

	// create the directory if it doesn't exist
	if !exists(jc.Dir) {
		err := jc.MkdirAll()
		if err != nil {
			return err
		}
	}

	// if file exists and saved is newer fail
	if exists(jc.Path()) {
		ondisk, err := NewFromFile(jc.Path())
		if err != nil {
			return err
		}
		if !ondisk.Saved.IsZero() && ondisk.Saved.After(jc.Updated) {
			return ErrorNewer
		}
	}

	// save it
	jc.Saved = time.Now()
	return os.WriteFile(jc.Path(), []byte(jc.String()), 0600)
}

// ForceSave is the same as Save but bypasses the latest Save check.
func (jc *Config) ForceSave() error {
	jc.mu.Lock()
	defer jc.mu.Unlock()

	// create the directory if it doesn't exist
	if !exists(jc.Dir) {
		err := jc.MkdirAll()
		if err != nil {
			return err
		}
	}

	// save it
	jc.Saved = time.Now()
	return os.WriteFile(jc.Path(), []byte(jc.String()), 0600)
}

// String implements the Stringer interface as a compressed (no indents)
// JSON string. Prints JSON with only ERROR if error occurred during
// marshaling.
func (jc Config) String() string {
	byt, err := json.Marshal(jc)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}

// PrettyPrint outputs the Data in an organized way for review (but
// without color).
func (jc Config) PrettyPrint() {
	_, max := jc.LongestKey()
	tpl := fmt.Sprintf("%%%vv %%v\n", max)
	keys := jc.Keys()
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf(tpl, k, jc.Data[k])
	}
}

// LongestKey returns the longest key in the data and its len value.
func (jc Config) LongestKey() (string, int) {
	ln := 0
	val := ""
	for k, v := range jc.Data {
		if len(k) > ln {
			ln = len(k)
			val = v
		}
	}
	return val, ln
}

// Keys returns an unordered list of keys from Data.
func (js Config) Keys() []string {
	keys := make([]string, len(js.Data))
	n := 0
	for k, _ := range js.Data {
		keys[n] = k
		n++
	}
	return keys
}

// MkdirAll attempts to create the directory and all parent directories
// for the Config.Dir (which should always be an absolute path). Called
// whenever Save() or ForceSave() are called.
func (c *Config) MkdirAll() error {
	return os.MkdirAll(c.Dir, 0700)
}
