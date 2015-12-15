package tilde

import (
	"os"
	"os/user"
	"strings"
	"testing"
)

var examples = []*struct {
	Example  string
	Expected string
}{
	{"", ""},
	{"~", "HOME"},
	{"hello_world", "hello_world"},
	{"/root/blah", "/root/blah"},
	{"~/foo/bar.txt", "HOME/foo/bar.txt"},
	{"~+/current/wd/", "PWD/current/wd"},
	{"~USER/zoink/narf", "HOME/zoink/narf"},
}

var currentUser *user.User
var currentWd string

func init() {
	var err error
	if currentUser, err = user.Current(); err != nil {
		panic(err)
	}
	if currentWd, err = os.Getwd(); err != nil {
		panic(err)
	}
}

func resolvePath(path string) string {
	path = strings.Replace(path, "HOME", currentUser.HomeDir, -1)
	path = strings.Replace(path, "PWD", currentWd, -1)
	path = strings.Replace(path, "USER", currentUser.Username, -1)
	return path
}

func TestTilde(t *testing.T) {
	for _, example := range examples {
		newPath, err := New(resolvePath(example.Example))
		if err != nil {
			t.Errorf("%s: Unable to resolve path", example.Example)
		}

		expectedPath := resolvePath(example.Expected)
		if newPath != expectedPath {
			t.Errorf("Expected path '%s', but got '%s'", expectedPath, newPath)
		}
	}
}
