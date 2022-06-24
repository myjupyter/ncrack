package ncrack

type HostOption func(*HostOptions)

type HostOptions struct {
	Protocol       arg
	Target         arg
	Port           arg
	ServiceOptions arg
}

// WithHostOption methods applies options to reffering host(s)
func WithHostOption(opts ...HostOption) Option {
	return func(c *Cracker) {
		ho := HostOptions{}
		for _, opt := range opts {
			opt(&ho)
		}
		c.args = append(c.args, ho.String())
	}
}

// WithHostOptionProtocols provides protocol
func WithHostOptionProtocols(p string) HostOption {
	return func(h *HostOptions) {
		h.Protocol.Val = p
	}
}

// WithHostOptionTarget add hosts that will be cracked
// ncrack supports hostnames, CIDR-addressing so you can use notation: "198.168.1.1/24",
// "10.0.0.1/28,10.0.0.2" and through octet range addressing: "192.168.0-250,255.1",
// "0-.0.0.1", but not both in the same host string
// To get more information see: https://nmap.org/ncrack/man.html
func WithHostOptionTarget(target string) HostOption {
	return func(h *HostOptions) {
		h.Target.Val = target
	}
}

// WithHostOptionPorts speccifies what port will be cracked
func WithHostOptionPorts(p uint16) HostOption {
	return func(h *HostOptions) {
		h.Port.Val = p
	}
}

// WithServiceOptions applies ServiceOptions to reffering host
// see service_options.go
func WithServiceOptions(opts ...ServiceOption) HostOption {
	return func(h *HostOptions) {
		so := ServiceOptions{}
		for _, opt := range opts {
			opt(&so)
		}
		h.ServiceOptions.Val = so
	}
}

func (h HostOptions) String() string {
	return h.Protocol.WithSchemaAfter() +
		h.Target.WithColonBetween(h.Port.WithColumnBetween(h.ServiceOptions.Arg()))
}
