// Package url contains utility for working with URL.
package url

import (
	"fmt"
	neturl "net/url"
)

// Build helps to build a clean URL
func Build(scheme, username, password, host string, port int) string {
	if port != 0 {
		host = fmt.Sprintf("%s:%d", host, port)
	}

	url := &neturl.URL{
		Scheme: scheme,
		Host:   host,
	}

	if username != "" {
		if password != "" {
			url.User = neturl.UserPassword(username, password)
		} else {
			url.User = neturl.User(username)
		}
	}

	return url.String()
}
