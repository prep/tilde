tilde
[![TravisCI](https://travis-ci.org/prep/tilde.svg?branch=master)](https://travis-ci.org/prep/tilde.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/prep/tilde)](https://goreportcard.com/report/github.com/prep/tilde)
[![GoDoc](https://godoc.org/github.com/prep/tilde?status.svg)](https://godoc.org/github.com/prep/tilde)
=====
This package provides tilde expansion functionality in Go.

Usage
-----
```go
import "github.com/prep/tilde"
```

```go
func main() {
  path, err := tilde.New("~/.bash_profile")
  if err != nil {
    // Handle error.
  }

  // path is now /home/user/.bash_profile.
}
```
