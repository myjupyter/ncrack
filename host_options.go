package ncrack

import "strconv"

type HostOption func(*HostOptions)

type HostOptions struct {
	Protocol       ProtocolType
	Target         string
	Port           uint16
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

func WithHostOptionProtocols(p ProtocolType) HostOption {
	return func(h *HostOptions) {
		h.Protocol = p
	}
}

func WithHostOptionTarget(target string) HostOption {
	return func(h *HostOptions) {
		h.Target = target
	}
}

func WithHostOptionPorts(p uint16) HostOption {
	return func(h *HostOptions) {
		h.Port = p
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
	if h.Protocol == ProtocolTypeUnspecified {
		return ""
	}
	if h.Target == "" {
		return ""
	}
	if h.Port == 0 {
		return ""
	}
	var services string
	if len(h.ServiceOptions) != 0 {
		services += "," + h.ServiceOptions.String()
	}
	return h.Protocol.String() + "://" + h.Target + ":" + strconv.Itoa(int(h.Port)) + services
}
