package conf

import (
	"os"
	"testing"
)

/*
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
*/

func TestConfigDir(t *testing.T) {
	t.Log(ConfigDir())
	// mocking the home directory is just not fucking worth it
	os.Setenv("XDG_CONFIG_HOME", "testdata/config")
	t.Log(ConfigDir())
}
