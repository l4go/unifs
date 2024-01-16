package unifs_test

import (
	"fmt"
	"io/fs"

	"github.com/l4go/unifs"
)

func ExampleNew() {
	var up unifs.UniPath
	var err error

	up, err = unifs.New("hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.New("/")
	fmt.Println(up, err == nil)
	up, err = unifs.New("/hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.New("/hoge/fuga/")
	fmt.Println(up, err == nil)
	up, err = unifs.New("/C:")
	fmt.Println(up, err == nil)
	up, err = unifs.New("/C:/hoge")
	fmt.Println(up, err == nil)

	// Output:
	//  false
	// / true
	// /hoge true
	// /hoge/fuga true
	// /C: true
	// /C:/hoge true
}

func ExampleNewFromFSPath() {
	var up unifs.UniPath
	var err error

	up, err = unifs.NewFromFSPath(".")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromFSPath("hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromFSPath("hoge/fuga")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromFSPath("/")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromFSPath("/hoge")
	fmt.Println(up, err == nil)
	up, err = unifs.NewFromFSPath("hoge/")
	fmt.Println(up, err == nil)

	// Output:
	// / true
	// /hoge true
	// /hoge/fuga true
	//  false
	//  false
	//  false
}

func ExampleUniPath_Join() {
	var err error
	var up, tp unifs.UniPath

	up, err = unifs.New("/test")
	fmt.Println(err == nil)
	tp, err = up.Join("sample.txt")
	fmt.Println(tp, err == nil)
	tp, err = up.Join("hoge/sample.txt")
	fmt.Println(tp, err == nil)
	tp, err = up.Join(`hoge\sample.txt`)
	fmt.Println(tp, err == nil)
	tp, err = up.Join("../sample.txt")
	fmt.Println(tp, err == nil)
	tp, err = up.Join("../../sample.txt")
	fmt.Println(tp, err == nil)
	// Output:
	// true
	// /test/sample.txt true
	// /test/hoge/sample.txt true
	//  false
	// /sample.txt true
	// /sample.txt true
}
func ExampleUniPath_Open() {
	var err error
	var up unifs.UniPath

	up, err = unifs.New("/test/sample.txt")
	fmt.Println(err == nil)
	_, err = up.Open(testfs)
	fmt.Println(err == nil)
	// Output:
	// true
	// true
}

func ExampleUniPath_Sub() {
	var err error
	var subfs fs.FS
	var up unifs.UniPath

	up, err = unifs.New("/test")
	fmt.Println(err == nil)
	subfs, err = up.Sub(testfs)
	fmt.Println(err == nil)
	_, err = unifs.Open(subfs, "/sample.txt")
	fmt.Println(err == nil)
	// Output:
	// true
	// true
	// true
}

func ExampleUniPath_Stat() {
	var err error
	var up unifs.UniPath

	up, err = unifs.New("/test/sample.txt")
	fmt.Println(err == nil)
	_, err = up.Stat(testfs)
	fmt.Println(err == nil)
	// Output:
	// true
	// true
}

func ExampleUniPath_ReadFile() {
	var err error
	var up unifs.UniPath

	up, err = unifs.New("/test/sample.txt")
	fmt.Println(err == nil)
	_, err = up.ReadFile(testfs)
	fmt.Println(err == nil)
	// Output:
	// true
	// true
}

func ExampleUniPath_ReadDir() {
	var err error
	var up unifs.UniPath

	up, err = unifs.New("/")
	fmt.Println(err == nil)
	_, err = up.ReadDir(testfs)
	fmt.Println(err == nil)
	up, err = unifs.New("/test")
	fmt.Println(err == nil)
	_, err = up.ReadDir(testfs)
	fmt.Println(err == nil)
	up, err = unifs.New("/hoge")
	fmt.Println(err == nil)
	_, err = up.ReadDir(testfs)
	fmt.Println(err == nil)
	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
}
