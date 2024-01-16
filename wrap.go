package unifs

import (
	"io/fs"
)

func Open(fsys fs.FS, name string) (fs.File, error) {
	up, err := ToFSPath(name)
	if err != nil {
		return nil, err
	}
	return fsys.Open(up)
}

func Sub(fsys fs.FS, name string) (fs.FS, error) {
	up, err := ToFSPath(name)
	if err != nil {
		return nil, err
	}
	return fs.Sub(fsys, up)
}

func Stat(fsys fs.FS, name string) (fs.FileInfo, error) {
	up, err := ToFSPath(name)
	if err != nil {
		return nil, err
	}
	return fs.Stat(fsys, up)
}

func Glob(fsys fs.FS, pattern string) ([]string, error) {
	up, err := ToFSPath(pattern)
	if err != nil {
		return nil, err
	}

	fn_lst, err := fs.Glob(fsys, up)
	if err != nil {
		return nil, err
	}

	up_lst := make([]string, len(fn_lst))
	for i, fn := range fn_lst {
		var e error
		up_lst[i], e = FromFSPath(fn)
		if e != nil {
			return nil, e
		}
	}

	return up_lst, nil
}

func ReadFile(fsys fs.FS, name string) ([]byte, error) {
	up, err := ToFSPath(name)
	if err != nil {
		return nil, err
	}
	return fs.ReadFile(fsys, up)
}

func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {
	up, err := ToFSPath(name)
	if err != nil {
		return nil, err
	}
	return fs.ReadDir(fsys, up)
}
