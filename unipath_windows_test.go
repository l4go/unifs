package unifs_test

import (
	"fmt"
	"os"

	"github.com/l4go/unifs"
)

func ExampleNewFromOSPath_win() {
	var up unifs.UniPath
	var err error

	// FIXME: Not tested
	os.Chdir(`C:\Users`)
	up, err = unifs.NewFromOSPath(".")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("hoge/fuga")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("C:")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromOSPath("C:/hoge")
	fmt.Println(up, err == nil)

	// Output:
	// /C:/Users true
	// /C:/Users/hoge true
	// /C:/Users/hoge/fuga true
	// /C: true
	// /C:/hoge true
}
