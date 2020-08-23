package vpn

import "net/url"

// WebSocketAddr ...
type WebSocketAddr struct {
	URL                   string
	InsecureSkipVerifyTLS bool
}

// Scheme ...
func (w WebSocketAddr) Scheme() string {
	u, err := url.Parse(w.URL)
	if err != nil {
		return ""
	}
	return u.Scheme
}

func (w WebSocketAddr) String() string {
	return w.URL
}
