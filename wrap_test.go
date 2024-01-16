package unifs_test

import (
	"fmt"
	"io/fs"

	"github.com/l4go/unifs"
)

func ExampleOpen() {
	var err error
	_, err = unifs.Open(testfs, "/test/sample.txt")
	fmt.Println(err == nil)
	// Output:
	// true
}

func ExampleSub() {
	var err error
	var subfs fs.FS

	subfs, err = unifs.Sub(testfs, "/test")
	fmt.Println(err == nil)
	_, err = unifs.Open(subfs, "/sample.txt")
	fmt.Println(err == nil)
	// Output:
	// true
	// true
}

func ExampleStat() {
	var err error
	_, err = unifs.Stat(testfs, "/test/sample.txt")
	fmt.Println(err == nil)
	// Output:
	// true
}

func ExampleReadFile() {
	var err error
	_, err = unifs.ReadFile(testfs, "/test/sample.txt")
	fmt.Println(err == nil)
	// Output:
	// true
}

func ExampleReadDir() {
	var err error
	_, err = unifs.ReadDir(testfs, "/")
	fmt.Println(err == nil)
	_, err = unifs.ReadDir(testfs, "/test")
	fmt.Println(err == nil)
	// Output:
	// true
	// true
}

func ExampleGlob() {
	ups, err := unifs.Glob(testfs, "/test/*.txt")
	if err != nil {
		return
	}
	for _, up := range ups {
		fmt.Println(up)
	}
	// Unordered output:
	// /test/sample.txt
	// /test/dummy.txt
}
