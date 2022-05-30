package ncrack

type ModuleOption func(*ModuleOptions)

type ModuleOptions struct {
	Protocol       ProtocolType
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

func WithModuleOptionServiceOptions(opts ...ServiceOption) ModuleOption {
	return func(m *ModuleOptions) {
		so := ServiceOptions{}
		for _, opt := range opts {
			opt(&so)
		}
		m.ServiceOptions = so
	}
}

func WithModuleOptionProtocol(p ProtocolType) ModuleOption {
	return func(m *ModuleOptions) {
		m.Protocol = p
	}
}

func (m ModuleOptions) String() string {
	if m.Protocol == ProtocolTypeUnspecified {
		return ""
	}
	services := m.ServiceOptions.String()
	if services == "" {
		return ""
	}
	return m.Protocol.String() + ":" + services

}
