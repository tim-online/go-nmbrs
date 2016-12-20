package url

import (
	nativeurl "net/url"
)

func MustParse(rawurl string) *nativeurl.URL {
	u, err := nativeurl.Parse(rawurl)
	if err != nil {
		panic(err)
	}

	return u
}
