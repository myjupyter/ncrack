package ncrack

import "strings"

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
		c.args = append(c.args, strings.Join(us, ","))
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
		c.args = append(c.args, "-")
	}
}
