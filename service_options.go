package ncrack

import (
	"strconv"
	"strings"
)

type ServiceOptions []string

type ServiceOption func(*ServiceOptions)

func (s ServiceOptions) String() string {
	return strings.Join([]string(s), ",")
}

// WithGlobalServiceOptions apply options to all hosts
func WithGlobalServiceOptions(opts ...ServiceOption) Option {
	return func(c *Cracker) {
		c.args = append(c.args, "-g")
		withServiceOptions(opts...)(c)
	}
}

func withServiceOptions(opts ...ServiceOption) Option {
	return func(c *Cracker) {
		so := make(ServiceOptions, 0, len(opts))
		for _, opt := range opts {
			opt(&so)
		}
		c.args = append(c.args, so.String())
	}
}

// WithServiceOptionSSL enable/disable SSL over service
func WithServiceOptionSSL(enable bool) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "ssl="+func() string {
			if enable {
				return "yes"
			}
			return "no"
		}())
	}
}

// WithServiceOptionPath provides path name for relative URLs
// The path is always relative to the hostname or IP address,
// so if you want to target something like http://foobar.com/login.php
// the path must take the value path=login.php
// For more info see https://nmap.org/ncrack/man.html
func WithServiceOptionPath(path string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "path="+path)
	}
}

// WithServiceOptionDB provides database name to ncrack. Some services
// like MongoDB require database name
func WithServiceOptionDB(dbname string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "db="+dbname)
	}
}

// WithServiceOptionDomain provides specific domain. Some services
// like WinRM require domain
func WithServiceOptionDomain(domain string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "domain="+domain)
	}
}

// WithServiceOptionMinConnLimit provides minimum connection limit
func WithServiceOptionMinConnLimit(min uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cl="+strconv.Itoa(int(min)))
	}
}

// WithServiceOptionMaxConnLimit provides maximum connection limit
func WithServiceOptionMaxConnLimit(max uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "CL="+strconv.Itoa(int(max)))
	}
}

// WithServiceOptionAuthTries provides maximum attempts for authentication tries
func WithServiceOptionAuthTries(tries uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "at="+strconv.Itoa(int(tries)))
	}
}

// WithServiceOptionConnDelay provides delay time between each new connection
// It's possible to specify timeout with "ms", "s", m" and "h" suffixes,
// seconds is default time unit
func WithServiceOptionConnDelay(delay string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cd="+delay)
	}
}

// WithServiceOptionConnRetries procvides maximum number for connection retries
func WithServiceOptionConnRetries(retries uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cr="+strconv.Itoa(int(retries)))
	}
}

// WithServiceOptionTimeout provides timout overall cracking process
// It's possible to specify timeout with "ms", "s", m" and "h" suffixes,
// seconds is default time unit
func WithServiceOptionTimeout(timeout string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "to="+timeout)
	}
}
