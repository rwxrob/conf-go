package conf_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/rwxrob/conf-go"
)

func ExampleNew_with_directory() {
	js := conf.New("testdata")
	fmt.Println(filepath.Base(js.Dir))
	// Output:
	// testdata
}

func ExampleNew_with_directory_and_file() {
	js := conf.New("testdata", "myown.json")
	fmt.Println(filepath.Base(js.Dir), js.File)
	// Output:
	// testdata myown.json
}

func ExampleConfigDir_xdg() {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	defer func() { os.Setenv("XDG_CONFIG_HOME", xdg) }()
	os.Setenv("XDG_CONFIG_HOME", "testdata")
	d := conf.ConfigDir()
	fmt.Println(d)
	// Output:
	// testdata/conf-go.test
}

func ExampleSet() {
	jc := conf.New()
	jc.Set("name", "Mr. Rob")
	fmt.Println(jc.Data)
	// Output:
	// {"name":"Mr. Rob"}
}

func ExampleGet() {
	jc := conf.New()
	jc.Set("name", "Mr. Rob")
	name := jc.Get("name")
	fmt.Println(name)
	// Output:
	// Mr. Rob
}

func ExampleGet_empty() {
	jc := conf.New()
	name := jc.Get("name")
	fmt.Printf("%q\n", name)
	// Output:
	// ""
}

func ExampleNewFromJSON() {
	jsn := `{"data":{"name": "Mr. Rob"}}`
	jc, _ := conf.NewFromJSON([]byte(jsn))
	name := jc.Get("name")
	fmt.Println(name)
	// Output:
	// Mr. Rob
}

func ExampleNewFromFile() {
	jc, _ := conf.NewFromFile("testdata/mapsample.json")
	name := jc.Get("name")
	fmt.Println(name)
	// Output:
	// Mr. Rob
}

func ExampleLoad() {
	jc := conf.New()

	jc.Set("Some", "Thing")
	fmt.Println(jc.Data)

	jc.Dir = "testdata"
	jc.File = "mapsample.json"

	jc.Load()
	fmt.Println(jc.Data)

	// Output:
	// {"Some":"Thing"}
	// {"name":"Mr. Rob"}
}

func ExampleSave() {
	T := new(testing.T)
	dir := T.TempDir()

	// initial config saved
	jc := conf.New(dir)
	jc.Set("name", "Mr. Rob")
	fmt.Println(jc.Save())

	// another process messes with it
	another := conf.New(dir)
	another.Saved.Add(2 * time.Minute)
	another.Set("another", "Mr. Rob")
	fmt.Println(another.ForceSave())

	// panics
	fmt.Println(jc.Save())

	// Output:
	// <nil>
	// <nil>
	// Newer config file detected.
}

func ExampleConfig_LongestKey() {
	jc := conf.New()
	jc.Set("short", "short")
	jc.Set("long", "long")
	jc.Set("reallylong", "reallylong")
	fmt.Println(jc.LongestKey())
	// Output:
	// reallylong 10
}

func ExampleConfig_Keys() {
	jc := conf.New()
	jc.Set("short", "short")
	jc.Set("long", "long")
	jc.Set("reallylong", "reallylong")
	for _, v := range jc.Keys() {
		fmt.Println(v)
	}
	// Unordered Output:
	// reallylong
	// short
	// long
}

func ExampleConfig_PrettyPrint() {
	jc := conf.New()
	jc.Set("short", "short")
	jc.Set("long", "long")
	jc.Set("reallylong", "reallylong")
	fmt.Println("------------")
	jc.PrettyPrint()
	// Output:
	// ------------
	//       long long
	// reallylong reallylong
	//      short short
}

/*
func BenchmarkSet(b *testing.B) {
	jc := conf.New()
	for n := 0; n < b.N; n++ {
		jc.Set("name", "Mr. Rob")
	}
}
*/
