package unifs

import (
	"io/fs"
	"path"
)

type UniPath struct {
	str string
}

var Zero = UniPath{str: ""}

func New(uni_name string) (UniPath, error) {
	if !ValidPath(uni_name) {
		return Zero, ErrInvalid
	}

	return UniPath{str: MustClean(uni_name)}, nil
}

func MustNew(uni_name string) UniPath {
	p, err := New(uni_name)
	if err != nil {
		panic("unifs: new: " + err.Error())
	}

	return p
}

func NewFromOSPath(os_name string) (UniPath, error) {
	up, err := FromOSPath(os_name)
	if err != nil {
		return UniPath{str: up}, err
	}

	return UniPath{str: up}, nil
}

func MustNewFromOSPath(os_path string) UniPath {
	p, err := NewFromOSPath(os_path)
	if err != nil {
		panic("unifs: new: " + err.Error())
	}

	return p
}

func NewFromFSPath(fs_name string) (UniPath, error) {
	up, err := FromFSPath(fs_name)
	if err != nil {
		return UniPath{str: up}, err
	}

	return UniPath{str: up}, nil
}

func MustNewFromFSPath(fs_name string) UniPath {
	p, err := NewFromFSPath(fs_name)
	if err != nil {
		panic("unifs: new: " + err.Error())
	}

	return p
}

func (up UniPath) panicIfZero() {
	if up.IsZero() {
		panic("unifs: join: Zero UniPath")
	}
}

func (up UniPath) IsZero() bool {
	return up.str == ""
}

func (up UniPath) String() string {
	return up.str
}

func (up UniPath) FSPath() string {
	up.panicIfZero()
	return toFSPath(up.str)
}

func (up UniPath) OSPath() string {
	up.panicIfZero()

	os_name, err := toOSPath(up.str)
	if err != nil {
		panic("unifs: ospath: " + err.Error())
	}

	return os_name
}

func (up UniPath) Join(names ...string) (UniPath, error) {
	up.panicIfZero()
	for _, n := range names {
		if !ValidSubPath(n) {
			return Zero, ErrInvalid
		}
	}
	j_lst := make([]string, len(names)+1)
	j_lst[0] = up.str
	copy(j_lst[1:], names)

	return UniPath{str: path.Join(j_lst...)}, nil
}

func (up UniPath) MustJoin(names ...string) UniPath {
	p, err := up.Join(names...)
	if err != nil {
		panic("unifs: join: " + err.Error())
	}

	return p
}

func (up UniPath) Open(fsys fs.FS) (fs.File, error) {
	up.panicIfZero()
	return Open(fsys, up.String())
}

func (up UniPath) Sub(fsys fs.FS) (fs.FS, error) {
	up.panicIfZero()
	return Sub(fsys, up.String())
}

func (up UniPath) Stat(fsys fs.FS) (fs.FileInfo, error) {
	up.panicIfZero()
	return Stat(fsys, up.String())
}

func (up UniPath) ReadFile(fsys fs.FS) ([]byte, error) {
	up.panicIfZero()
	return ReadFile(fsys, up.String())
}

func (up UniPath) ReadDir(fsys fs.FS) ([]fs.DirEntry, error) {
	up.panicIfZero()
	return ReadDir(fsys, up.String())
}

func (up UniPath) Glob(fsys fs.FS) ([]UniPath, error) {
	up.panicIfZero()
	ups, err := Glob(fsys, up.String())
	if err != nil {
		return nil, err
	}

	ret := make([]UniPath, len(ups))
	for i, u := range ups {
		var e error
		ret[i], e = New(u)
		if e != nil {
			return nil, e
		}
	}

	return ret, nil
}
