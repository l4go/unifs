//go:build !windows

package unifs_test

import (
	"fmt"

	"github.com/l4go/unifs"
)

func ExampleToOSPath_unix() {
	var name string
	var err error
	name, err = unifs.ToOSPath("/C:")
	fmt.Println(err == nil, name)
	name, err = unifs.ToOSPath("/C:/")
	fmt.Println(err == nil, name)
	name, err = unifs.ToOSPath("/C:/hoge")
	fmt.Println(err == nil, name)
	name, err = unifs.ToOSPath("/C:/hoge/")
	fmt.Println(err == nil, name)
	// Output:
	// true /C:
	// true /C:/
	// true /C:/hoge
	// true /C:/hoge/
}
