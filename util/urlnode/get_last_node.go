package urlnode

import (
	"golang.org/x/xerrors"
	liburl "net/url"
	"strings"
)

func GetLastNodeFromString(url string) (node string, err error) {
	parsed, err := liburl.Parse(url)
	if err != nil {
		return "", xerrors.Errorf("Can't parse url.URL from URL(%s): %w", url, err)
	}

	return GetLastNode(parsed)
}

func GetLastNode(url *liburl.URL) (node string, err error) {
	urlPath := url.Path
	pathFragments := strings.Split(urlPath, "/")

	return pathFragments[len(pathFragments)-1], nil
}
