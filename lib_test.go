package conf

import (
	"testing"
)

func TestExists(t *testing.T) {
	if !exists("testdata") || exists("tstdata") {
		t.Fail()
	}
}

func TestNew_relative_dir(t *testing.T) {
	jc := New("testdata")
	t.Log(jc.Dir)
}

func TestConfigDir(t *testing.T) {
	t.Log(ConfigDir())
}

func TestConfigPath(t *testing.T) {
	c := New("testdata", "pathtest")
	t.Log(c.Path())
	t.Log("")
}

func TestNewFromFile_missing(t *testing.T) {
	jc, err := NewFromFile("testdata/mpsample.json")
	t.Log(jc, err)
}

func TestLoad_missing(t *testing.T) {
	jc := New()
	jc.Set("Some", "Thing")
	t.Log(jc.Data)
	jc.Dir = "tetdata"
	jc.File = "mapsample.json"
	err := jc.Load()
	if err == nil {
		t.Fail()
	}
}
