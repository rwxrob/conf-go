package conf_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/conf-go"
)

func ExampleHome() {
	os.Setenv("HOME", "/home/rwxrob")
	fmt.Println(conf.Home())

	// Output:
	// /home/rwxrob/.config
}

func ExampleExeDir() {
	os.Setenv("HOME", "/home/rwxrob")
	fmt.Println(conf.ExeDir())

	// Output:
	// /home/rwxrob/.config/conf-go.test
}

func ExampleExeDirFile() {
	os.Setenv("HOME", "/home/rwxrob")
	fmt.Println(conf.ExeDirFile("values"))

	// Output:
	// /home/rwxrob/.config/conf-go.test/values
}

func ExampleNewMap() {
	m := conf.NewMap()
	m.Set("foo", "FOO")
	fmt.Println(m.Get("foo"))
	m.Print()
	fmt.Println(m.JSON())
	m.PrintJSON()

	// Output:
	// FOO
	// foo=FOO
	// {"foo":"FOO"}
	// {"foo":"FOO"}
}

func ExampleKeys() {
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	fmt.Println(m.Keys())

	// Output:
	// [bar foo]
}

func ExampleMap_Save() {
	os.Setenv("HOME", "testdata")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Save()
	buf, _ := os.ReadFile("testdata/.config/conf-go.test/values")
	fmt.Println(string(buf))

	// Output:
	// foo=FOO
}
