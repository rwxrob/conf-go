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
