package conf

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rv := m.Run()
	os.RemoveAll("./testdata/.config")
	os.Exit(rv)
}
