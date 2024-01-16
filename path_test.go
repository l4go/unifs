package unifs_test

import (
	"fmt"

	"github.com/l4go/unifs"
)

func ExampleClean() {
	fmt.Println(unifs.MustClean("/C:/foo/../bar/"))
	fmt.Println(unifs.MustClean("/foo/../bar/"))
	fmt.Println(unifs.MustClean("/foo/../../bar/"))

	// Output:
	// /C:/bar
	// /bar
	// /bar
}

func ExampleJoin() {
	fmt.Println(unifs.MustJoin("/C:/foo/", "../bar"))
	fmt.Println(unifs.MustJoin("/foo/", "../bar/"))
	fmt.Println(unifs.MustJoin("/foo/", "../../bar"))

	// Output:
	// /C:/bar
	// /bar
	// /bar
}

func ExampleValidPath() {
	fmt.Println(unifs.ValidPath("/C:/"))
	fmt.Println(unifs.ValidPath("/C:/hoge"))
	fmt.Println(unifs.ValidPath("/"))
	fmt.Println(unifs.ValidPath(""))
	fmt.Println(unifs.ValidPath("hoge/fuga"))
	fmt.Println(unifs.ValidPath("../hoge"))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleValidSubPath() {
	fmt.Println(unifs.ValidSubPath("hoge"))
	fmt.Println(unifs.ValidSubPath("hoge/fuga"))
	fmt.Println(unifs.ValidSubPath("../hoge"))
	fmt.Println(unifs.ValidSubPath(`hoge\fuga`))
	fmt.Println(unifs.ValidSubPath("hoge\x00"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleToFSPath() {
	fmt.Println(unifs.ToFSPath("/C:"))
	fmt.Println(unifs.ToFSPath("/C:/"))
	fmt.Println(unifs.ToFSPath("/C:/hoge"))
	fmt.Println(unifs.ToFSPath("/C:/hoge/"))

	// Output:
	// C: <nil>
	// C: <nil>
	// C:/hoge <nil>
	// C:/hoge <nil>
}

func ExampleFromFSPath() {
	var name string
	var err error
	name, err = unifs.FromFSPath(".")
	fmt.Println(name, err == nil)
	name, err = unifs.FromFSPath("hoge")
	fmt.Println(name, err == nil)
	name, err = unifs.FromFSPath("hoge/fuga")
	fmt.Println(name, err == nil)
	name, err = unifs.FromFSPath("C:")
	fmt.Println(name, err == nil)
	name, err = unifs.FromFSPath("C:/hoge")
	fmt.Println(name, err == nil)

	// Output:
	// / true
	// /hoge true
	// /hoge/fuga true
	// /C: true
	// /C:/hoge true
}
