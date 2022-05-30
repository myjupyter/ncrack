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

func WithServiceOptionPath(path string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "path="+path)
	}
}

func WithServiceOptionDB(dbname string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "db="+dbname)
	}
}

func WithServiceOptionDomain(domain string) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "domain="+domain)
	}
}

func WithServiceOptionMinConnLimit(min uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cl="+strconv.Itoa(int(min)))
	}
}

func WithServiceOptionMaxConnLimit(max uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "CL="+strconv.Itoa(int(max)))
	}
}

func WithServiceOptionAuthTries(tries uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "at="+strconv.Itoa(int(tries)))
	}
}

func WithServiceOptionConnDelay(delay uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cd="+strconv.Itoa(int(delay)))
	}
}

func WithServiceOptionConnRetries(retries uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "cr="+strconv.Itoa(int(retries)))
	}
}

func WithServiceOptionTimeout(timeout uint32) ServiceOption {
	return func(s *ServiceOptions) {
		*s = append(*s, "to="+strconv.Itoa(int(timeout)))
	}
}
