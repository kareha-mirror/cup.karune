package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Get(u *url.URL) (string, error) {
	switch u.Scheme {
	case "http", "https":
		return getHTTP(u)

	case "file":
		return getFile(u)

	default:
		return "", fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
}

func getHTTP(u *url.URL) (string, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getFile(u *url.URL) (string, error) {
	f, err := os.Open(u.Path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
