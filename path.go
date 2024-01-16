package unifs

import (
	"io/fs"
	"path"
	"strings"
)

var invalid_path_chats = "\\\x00"

func MustClean(uni_name string) string {
	p, err := Clean(uni_name)
	if err != nil {
		panic("unifs: clean: " + err.Error())
	}

	return p
}

func Clean(uni_name string) (string, error) {
	if !ValidPath(uni_name) {
		return "", ErrInvalid
	}

	return path.Clean(uni_name), nil
}

func MustJoin(names ...string) string {
	p, err := Join(names...)
	if err != nil {
		panic("unifs: join: " + err.Error())
	}

	return p
}

func Join(names ...string) (string, error) {
	if !ValidPath(names[0]) {
		return "", ErrInvalid
	}

	for _, n := range names[1:] {
		if !ValidSubPath(n) {
			return "", ErrInvalid
		}
	}

	return path.Join(names...), nil
}

func ValidSubPath(uni_name string) bool {
	return !strings.ContainsAny(uni_name, invalid_path_chats)
}

func ValidPath(uni_name string) bool {
	return ValidSubPath(uni_name) && path.IsAbs(uni_name)
}

func toFSPath(uni_name string) string {
	if uni_name == "/" {
		return "."
	}

	return strings.Trim(uni_name, "/")
}

func ToFSPath(uni_name string) (string, error) {
	var err error
	uni_name, err = Clean(uni_name)
	if err != nil {
		return "", err
	}

	return toFSPath(uni_name), nil
}

func ToFSPaths(uni_names []string) ([]string, error) {
	res := make([]string, len(uni_names))
	var err error
	for i, n := range uni_names {
		res[i], err = ToFSPath(n)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func ToOSPath(uni_name string) (string, error) {
	return toOSPath(uni_name)
}

func ToOSPaths(uni_names []string) ([]string, error) {
	res := make([]string, len(uni_names))
	var err error
	for i, n := range uni_names {
		res[i], err = ToOSPath(n)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func FromFSPath(fs_name string) (string, error) {
	if !fs.ValidPath(fs_name) {
		return "", ErrInvalid
	}

	if fs_name == "." {
		return "/", nil
	}

	return path.Clean(path.Join("/", fs_name)), nil
}

func FromOSPath(os_name string) (string, error) {
	return fromOSPath(os_name)
}
