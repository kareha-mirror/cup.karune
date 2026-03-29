package util

import (
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

var schemeRe = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9+.-]*://")

func ResolvePath(s string) (*url.URL, error) {
	if schemeRe.MatchString(s) {
		return url.Parse(s)
	}
	_, err := os.Stat(s)
	if err == nil {
		absPath, err := filepath.Abs(s)
		if err != nil {
			panic(err)
		}
		u := &url.URL{Scheme: "file"}
		u.Path = filepath.ToSlash(absPath)
		return u, nil
	}
	return url.Parse("https://" + s)
}
