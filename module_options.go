package ncrack

import "strings"

type ModuleOption func(*ModuleOptions)

type ModuleOptions struct {
	Protocol       string
	ServiceOptions ServiceOptions
}

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

func WithModuleServiceOptions(opts ...ServiceOption) ModuleOption {
	return func(m *ModuleOptions) {
		so := ServiceOptions{}
		for _, opt := range opts {
			opt(&so)
		}
		m.ServiceOptions = so
	}
}

func WithModuleOptionProtocol(p ...string) ModuleOption {
	return func(m *ModuleOptions) {
		m.Protocol = strings.Join(p, ",")
	}
}

func (m ModuleOptions) String() string {
	if m.Protocol == "" {
		return ""
	}
	services := m.ServiceOptions.String()
	if services == "" {
		return ""
	}
	return m.Protocol + ":" + services

}
