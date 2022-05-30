package ncrack

import "strconv"

type HostOption func(*HostOptions)

type HostOptions struct {
	Protocols      []ProtocolType
	Target         string
	Ports          []uint16
	ServiceOptions ServiceOptions
}

func WithHostOption(opts ...HostOption) Option {
	return func(c *Cracker) {
		ho := HostOptions{}
		for _, opt := range opts {
			opt(&ho)
		}
		c.args = append(c.args, ho.String())
	}
}

func WithHostOptionProtocols(ps ...ProtocolType) HostOption {
	return func(h *HostOptions) {
		h.Protocols = append(h.Protocols, ps...)
	}
}

func WithHostOptionTarget(target string) HostOption {
	return func(h *HostOptions) {
		h.Target = target
	}
}

func WithHostOptionPorts(ps ...uint16) HostOption {
	return func(h *HostOptions) {
		h.Ports = append(h.Ports, ps...)
	}
}

func WithHostOptionServiceOptions(opts ...ServiceOption) HostOption {
	return func(h *HostOptions) {
		if h.ServiceOptions == nil {
			h.ServiceOptions = make(ServiceOptions, 1)
		}
		for _, opt := range opts {
			opt(&h.ServiceOptions)
		}
	}
}

func (h HostOptions) String() string {
	var protocols string
	for i := range h.Protocols {
		protocols += h.Protocols[i].String()
	}
	if protocols != "" {
		protocols += "://"
	}
	var ports string
	for i := range h.Ports {
		if i > 0 {
			ports += ","
		}
		ports += strconv.Itoa(int(h.Ports[i]))
	}
	if ports != "" {
		ports = ":" + ports
	}
	services := h.ServiceOptions.String()

	return protocols + h.Target + ports + services
}
