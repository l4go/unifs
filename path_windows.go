package unifs

import (
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var absPathRe = regexp.MustCompile(`^[a-zA-Z]:`)
var drivePathRe = regexp.MustCompile(`^[a-zA-Z]:$`)

func absOSPath(os_name string) (string, error) {
	if strings.HasPrefix(os_name, `\`) {
		return "", ErrInvalid
	}

	if !absPathRe.MatchString(os_name) {
		abs_name, err := filepath.Abs(os_name)
		if err != nil {
			return "", err
		}

		os_name = abs_name
	}
	if drivePathRe.MatchString(os_name) {
		os_name += `\`
	}

	return os_name, nil
}

func toOSPath(uni_name string) (string, error) {
	if len(uni_name) < 1 || uni_name[0] != '/' {
		return "", ErrInvalid
	}
	os_name := strings.ReplaceAll(uni_name[1:], "/", `\`)

	if !absPathRe.MatchString(os_name) {
		return "", ErrInvalid
	}
	if drivePathRe.MatchString(os_name) {
		os_name += `\`
	}

	return os_name, nil
}

func fromOSPath(os_name string) (string, error) {
	abs_name, err := absOSPath(os_name)
	if err != nil {
		return "", err
	}
	os_name = abs_name

	uni_name := "/" + strings.ReplaceAll(os_name, `\`, "/")

	return path.Clean(uni_name), nil
}
