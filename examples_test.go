package conf_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/conf-go"
)

// -------------------------- base functions --------------------------

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

// --------------------------- Map interface --------------------------

func ExampleMap_Keys() {
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	fmt.Println(m.Keys())
	// Output:
	// [bar foo]
}

func ExampleMap_String() {
	m := conf.NewMap()
	m.Set("foo", "FOO")
	fmt.Println(m)
	// Output:
	// foo=FOO
}

func ExampleMap_Name() {
	m := conf.NewMap()
	fmt.Println("Default executable name: " + m.Name())
	m.SetName("foo")
	fmt.Println(m.Name())
	// Output:
	// Default executable name: conf-go.test
	// foo
}

// ----------------------- return *mapStruct/Map ----------------------

func ExampleParse() {
	m, err := conf.Parse([]byte("foo=FOO\r\nbar=BAR\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("bar"))
	fmt.Println(m.Raw())
	fmt.Println(m)
	// Output:
	// FOO
	// BAR
	// map[bar:BAR foo:FOO]
	// bar=BAR
	// foo=FOO
}

func ExampleParse_error_NoEqualSign() {
	_, err := conf.Parse([]byte("foo FOO\n"))
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// invalid config (line 1): foo FOO
}

func ExampleParse_unexpected_Key() {
	m, err := conf.Parse([]byte("foo =FOO\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", m.Get("foo"))
	fmt.Println(m.Raw())

	// Output:
	// ""
	// map[foo :FOO]
}

func ExampleParse_unexpected_Val() {
	m, err := conf.Parse([]byte("foo= FOO\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", m.Get("foo"))
	fmt.Println(m.Raw())
	// Output:
	// " FOO"
	// map[foo: FOO]
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

func ExampleWrite() {
	os.Setenv("HOME", "testdata")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	conf.Write(m)
	buf, err := os.ReadFile("testdata/.config/conf-go.test/values")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
	// Output:
	// bar=BAR
	// foo=FOO
}

func ExampleRead() {
	os.Setenv("HOME", "testdata")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	m.Set("other", "one")
	conf.Write(m)
	newm, err := conf.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newm)
	// Output:
	// bar=BAR
	// foo=FOO
	// other=one
}
