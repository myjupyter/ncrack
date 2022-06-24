package ncrack

import "strings"

type ModuleOption func(*ModuleOptions)

type ModuleOptions struct {
	Protocol       arg
	ServiceOptions arg
}

// Options in this mode apply to all hosts that are associated with the particular service/module
func WithModuleOption(opts ...ModuleOption) Option {
	return func(c *Cracker) {
		mo := ModuleOptions{}
		for _, opt := range opts {
			opt(&mo)
		}
		c.args = append(c.args, "-m")
		c.args = append(c.args, mo.String())
	}
}

// WithModuleServiceOptions provides ServiceModule to reffering module
// see service_options.go
func WithModuleServiceOptions(opts ...ServiceOption) ModuleOption {
	return func(m *ModuleOptions) {
		so := ServiceOptions{}
		for _, opt := range opts {
			opt(&so)
		}
		m.ServiceOptions.Val = so
	}
}

// WithModuleOptionProtocol defines module type
func WithModuleOptionProtocol(p ...string) ModuleOption {
	return func(m *ModuleOptions) {
		m.Protocol.Val = strings.Join(p, ",")
	}
}

func (m ModuleOptions) String() string {
	return m.Protocol.WithColonAfter() + m.ServiceOptions.Arg()
}
