package tilde

import (
	"os"
	"os/user"
	"path"
	"regexp"
)

var pathRegex = regexp.MustCompile(`(~)([^/]*)(/?.*)`)

// New takes a path that starts with a tilde and expands it to its full path.
func New(p string) (string, error) {
	if len(p) < 1 || p[0] != '~' {
		return p, nil
	}

	results := pathRegex.FindStringSubmatch(p)[2:]

	var tildePath string

	switch results[0] {
	case "":
		u, err := user.Current()
		if err != nil {
			return "", err
		}

		tildePath = u.HomeDir
	case "+":
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		tildePath = pwd
	default:
		u, err := user.Lookup(results[0])
		if err != nil {
			return "", err
		}

		tildePath = u.HomeDir
	}

	return path.Join(tildePath, results[1]), nil
}
