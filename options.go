package ncrack

import (
	"strings"
)

type Cracker struct {
	args []string
}

func (c Cracker) Args() []string {
	return c.args
}

type Option func(*Cracker)

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
func WithHosts(hs []string) Option {
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

func WithUser(us []string) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "--user")
		c.args = append(c.args, us...)
	}
}

func WithPass(ps []string) Option {
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
