package ncrack

import "strconv"

type HostOption func(*HostOptions)

type HostOptions struct {
	Protocol       string
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

func WithHostOptionProtocols(p string) HostOption {
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
	if h.Protocol == "" {
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
	return string(h.Protocol) + "://" + h.Target + ":" + strconv.Itoa(int(h.Port)) + services
}
