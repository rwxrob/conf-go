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

func ExampleExe() {
	os.Setenv("HOME", "/home/rwxrob")
	fmt.Println(conf.Exe())

	// Output:
	// /home/rwxrob/.config/conf-go.test
}

func ExampleExeFile() {
	os.Setenv("HOME", "/home/rwxrob")
	fmt.Println(conf.ExeFile("values"))

	// Output:
	// /home/rwxrob/.config/conf-go.test/values
}

func ExampleMapStruct() {
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
