package conf

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

type Map interface {
	Get(key string) string
	Set(key string, val string) error
	Raw(m map[string]string)
	Save() error
}

type MapStruct struct {
	sync.RWMutex
	m map[string]string
}

func NewMap() *MapStruct {
	m := new(MapStruct)
	m.m = map[string]string{}
	return m
}

// Raw returns the inner map[string]string for direct manipulation that
// bypasses locking such as when ranging or sorting.
func (m *MapStruct) Raw() map[string]string { return m.m }

func (m *MapStruct) Get(key string) string {
	m.RLock()
	defer m.RUnlock()
	return m.m[key]
}

// Set sets the value of key to the string. Warning: use of line returns
// will *not* be escaped automatically and will break the Save and Load
// functions. This is to preserve the performance of a single character
// delimiter per line without any additional overhead.
func (m *MapStruct) Set(key, val string) error {
	m.RLock()
	defer m.RUnlock()
	m.m[key] = val
	// TODO add line return escapes
	return m.Save()
}

func (m *MapStruct) Save() error {
	// TODO
	return nil
}

func (m *MapStruct) Keys() []string {
	keys := make([]string, len(m.m))
	n := 0
	for k, _ := range m.m {
		keys[n] = k
		n++
	}
	sort.Strings(keys)
	return keys
}

// String implements the fmt.Stringer interface.
func (m MapStruct) String() string {
	buf := ""
	for _, k := range m.Keys() {
		buf = buf + fmt.Sprintf("%v=%v\n", k, m.m[k])
	}
	return buf
}

func (m MapStruct) JSON() string {
	byt, err := json.Marshal(m.m)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":%q", err)
	}
	return string(byt)
}

func (m MapStruct) Print()     { fmt.Print(m) }
func (m MapStruct) PrintJSON() { fmt.Println(m.JSON()) }
