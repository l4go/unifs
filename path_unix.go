//go:build !windows

package unifs

import (
	"path"
	"path/filepath"
)

func absOSPath(os_name string) (string, error) {
	if len(os_name) > 0 && os_name[0] != '/' {
		abs_name, err := filepath.Abs(os_name)
		if err != nil {
			return "", err
		}

		os_name = abs_name
	}

	return os_name, nil
}

func toOSPath(uni_name string) (string, error) {
	if len(uni_name) < 1 || uni_name[0] != '/' {
		return "", ErrInvalid
	}

	return uni_name, nil
}

func fromOSPath(os_name string) (string, error) {
	abs_name, err := absOSPath(os_name)
	if err != nil {
		return "", err
	}

	return path.Clean(abs_name), nil
}
