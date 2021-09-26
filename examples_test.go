package conf_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/conf-go"
)

func ExampleNewMap() {
	os.Setenv("HOME", "/home/rwxrob")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	fmt.Println(m.Get("foo"))
	m.Print()
	fmt.Println(m.JSON())
	m.PrintJSON()
	fmt.Println(m.Home())
	fmt.Println(m.Name())
	fmt.Println(m.File())
	// Output:
	// FOO
	// foo=FOO
	// {"foo":"FOO"}
	// {"foo":"FOO"}
	// /home/rwxrob/.config
	// conf-go.test
	// values
}

func ExampleNewMap_xdg_config_home() {
	os.Setenv("XDG_CONFIG_HOME", "/tmp/config")
	m := conf.NewMap()
	fmt.Println(m.Path())
	os.Unsetenv("XDG_CONFIG_HOME") // so other tests will work
	// Output:
	// /tmp/config/conf-go.test/values
}

func ExampleMap_Raw() {
	m := conf.NewMap()
	fmt.Println(m.Raw())
	m.Set("foo", "FOO")
	fmt.Println(m.Raw())
	// Output:
	// map[]
	// map[foo:FOO]
}

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

func ExampleMap_Home() {
	os.Setenv("HOME", "/home/rwxrob")
	m := conf.NewMap()
	fmt.Println("Default config home: " + m.Home())
	os.Setenv("XDG_CONFIG_HOME", "/tmp/myxdg_config")
	fmt.Println("Cached home: " + m.Home())
	m.SetHome("/tmp/config")
	fmt.Println("Changed home: " + m.Home())
	os.Unsetenv("XDG_CONFIG_HOME") // so other tests will work
	// Output:
	// Default config home: /home/rwxrob/.config
	// Cached home: /home/rwxrob/.config
	// Changed home: /tmp/config
}

func ExampleMap_Name() {
	m := conf.NewMap()
	fmt.Println("Default subdirectory name: " + m.Name())
	m.SetName("foo")
	fmt.Println(m.Name())
	// Output:
	// Default subdirectory name: conf-go.test
	// foo
}

func ExampleMap_File() {
	m := conf.NewMap()
	fmt.Println("Default file: " + m.File())
	m.SetFile("other")
	fmt.Println(m.File())
	// Output:
	// Default file: values
	// other
}

func ExampleMap_Path() {
	os.Setenv("HOME", "/home/rwxrob")
	m := conf.NewMap()
	fmt.Println(m.Path())
	// Output:
	// /home/rwxrob/.config/conf-go.test/values
}

func ExampleParse() {
	m := conf.NewMap()
	err := m.Parse([]byte("foo=FOO\r\nbar=BAR\n"))
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
	m := conf.NewMap()
	err := m.Parse([]byte("foo FOO\n"))
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// invalid config (line 1): foo FOO
}

func ExampleParse_unexpected_Key() {
	m := conf.NewMap()
	err := m.Parse([]byte("foo =FOO\n"))
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
	m := conf.NewMap()
	err := m.Parse([]byte("foo= FOO\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", m.Get("foo"))
	fmt.Println(m.Raw())
	// Output:
	// " FOO"
	// map[foo: FOO]
}

func ExampleWrite() {
	os.Setenv("HOME", "testdata")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	m.Write()
	buf, err := os.ReadFile("testdata/.config/conf-go.test/values")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
	// Output:
	// bar=BAR
	// foo=FOO
}

func ExampleWrite_bork() {
	os.Setenv("HOME", "/bork")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	err := m.Write()
	fmt.Println(err)
	// Output:
	// mkdir /bork: permission denied
}

func ExampleRead() {
	os.Setenv("HOME", "testdata")
	m := conf.NewMap()
	m.Set("foo", "FOO")
	m.Set("bar", "BAR")
	m.Set("other", "one")
	m.Write()
	err := m.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	// Output:
	// bar=BAR
	// foo=FOO
	// other=one
}

func ExampleRead_bork() {
	os.Setenv("HOME", "bork")
	m := conf.NewMap()
	err := m.Read()
	fmt.Println(err)
	// Output:
	// open bork/.config/conf-go.test/values: no such file or directory
}

func ExampleEscape() {
	fmt.Printf("%q\n", conf.Escape("some\nthing"))
	fmt.Printf("%q\n", conf.Escape("foo\r\nbar\n"))
	// Output:
	// "some\\nthing"
	// "foo\\r\\nbar\\n"
}
