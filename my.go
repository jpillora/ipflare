package ipflare

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// My IP address, according to Cloudflare.
// Can be either IPv4 or IPv6.
func My() (net.IP, error) {
	resp, err := http.Get("https://cloudflare.com/cdn-cgi/trace")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(b), "\n") {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		k := kv[0]
		v := kv[1]
		if k != "ip" {
			continue
		}
		ip := net.ParseIP(v)
		if ip == nil {
			return nil, fmt.Errorf("my ip invalid (%s)", v)
		}
		return ip, nil
	}
	return nil, errors.New("my ip not found on page")
}
