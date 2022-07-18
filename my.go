package ipflare

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
)

var re = regexp.MustCompile(`id="cf-footer-ip">([a-f0-9\.:]+)</span>`)

//My IP address, according to Cloudflare.
//Can be either IPv4 or IPv6.
func My() (net.IP, error) {
	resp, err := http.Get("https://www.cloudflare.com/learning/dns/glossary/what-is-my-ip-address/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m := re.FindSubmatch(b)
	if len(m) == 0 {
		return nil, errors.New("my ip not found on page")
	}
	s := string(m[1])
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, fmt.Errorf("my ip invalid (%s)", s)
	}
	return ip, nil
}
