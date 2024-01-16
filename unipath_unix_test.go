//go:build !windows

package unifs_test

import (
	"fmt"
	"os"

	"github.com/l4go/unifs"
)

func ExampleNewFromOSPath_unix() {
	var up unifs.UniPath
	var err error

	os.Chdir("/usr")
	up, err = unifs.NewFromOSPath(".")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("hoge/fuga")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("/")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("/hoge")
	fmt.Println(up, err == nil)

	// Output:
	// /usr true
	// /usr/hoge true
	// /usr/hoge/fuga true
	// / true
	// /hoge true
}
