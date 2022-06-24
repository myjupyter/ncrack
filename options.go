package ncrack

import (
	"context"
	"strings"
)

type Option func(*Cracker)

// WithContext pass context to ncrack
func WithContext(ctx context.Context) Option {
	return func(c *Cracker) {
		c.ctx = ctx
	}
}

// WithNmapXML allows to use input from Nmap's -oX XML output format
func WithNmapXML(path string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-iX", path)
	}
}

// WithNmapNormal allows to use input from Nmap's -oN Normal output format
func WithNmapNormal(path string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-iN", path)
	}
}

// WithHosts add hosts that will be cracked
// ncrack supports hostnames, CIDR-addressing so you can use notation: "198.168.1.1/24",
// "10.0.0.1/28,10.0.0.2" and through octet range addressing: "192.168.0-250,255.1",
// "0-.0.0.1", but not both in the same host string
// To get more information see: https://nmap.org/ncrack/man.html
func WithHosts(hs ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, hs...)
	}
}

// WithHostsExclude excludes hosts from host range,
// supports hostnames, CIDR-addressing, octet range addressing
func WithHostsExclude(hs []string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--exclude")
		c.args = append(c.args, hs...)
	}
}

// WithIPv6 enable to crack hosts by IPv6
func WithIPv6() Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-6")
	}
}

// WithUser puts username list through command line
func WithUser(us ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--user")
		c.args = append(c.args, strings.Join(us, ","))
	}
}

// WithPass puts password list through command line
func WithPass(ps ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--pass")
		c.args = append(c.args, strings.Join(ps, ","))
	}
}

// WithPasswordsFirst makes ncrack iterate the password list for each username,
// by default ncrack iterate the username list for each passowrd
func WithPasswordsFirst() Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--passwords-first")
	}
}

// WithProxy option makes ncrack perform cracking through proxy
// See: https://nmap.org/ncrack/man.html
func WithProxy(addr string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--proxy")
		c.args = append(c.args, addr)
	}
}

// WithXMLOutput option makes ncrack output XML output to stdout
func WithXMLOutput() Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-oX")
		c.args = append(c.args, "-")
	}
}

// WithFilter provides function for service filtration
func WithFilter(fn func(*Service) bool) Option {
	return func(c *Cracker) {
		c.filterFunc = fn
	}
}
