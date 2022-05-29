package ncrack

import (
	"strconv"
	"strings"
)

type Option func(*Cracker)

// Input from Nmap's -oX XML output format
func WithNmapXML(path string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-iX", path)
	}
}

func WithNmapNormal(path string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-iN", path)
	}
}

func WithHosts(hs ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, hs...)
	}
}

func WithHostsExclude(hs []string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--exclude")
		c.args = append(c.args, hs...)
	}
}

func WithUser(us ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--user")
		c.args = append(c.args, us...)
	}
}

func WithPass(ps ...string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--pass")
		c.args = append(c.args, strings.Join(ps, ","))
	}
}

func WithPasswordsFirst() Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--passwords-first")
	}
}

func WithProxy(addr string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--proxy")
		c.args = append(c.args, addr)
	}
}

func WithXMLOutput() Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-oX")
	}
}

type HostOption struct {
	Protocols []ProtocolType
	Target    string
	Ports     []uint16
}

func (h HostOption) String() string {
	var protocols string
	for i := range h.Protocols {
		protocols += h.Protocols[i].String()
	}
	var ports string
	for i := range h.Ports {
		ports += strconv.Itoa(int(h.Ports[i]))
	}
	return protocols + "://" + h.Target + ":" + ports
}

func WithPerHostOption(hs ...HostOption) Option {
	return func(c *Cracker) {
		for i := range hs {
			c.args = append(c.args, hs[i].String())
		}
	}
}
