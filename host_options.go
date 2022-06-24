package ncrack

import "strconv"

type HostOption func(*HostOptions)

type HostOptions struct {
	Protocol       string
	Target         string
	Port           uint16
	ServiceOptions ServiceOptions
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
		h.Protocol = p
	}
}

// WithHostOptionTarget add hosts that will be cracked
// ncrack supports hostnames, CIDR-addressing so you can use notation: "198.168.1.1/24",
// "10.0.0.1/28,10.0.0.2" and through octet range addressing: "192.168.0-250,255.1",
// "0-.0.0.1", but not both in the same host string
// To get more information see: https://nmap.org/ncrack/man.html
func WithHostOptionTarget(target string) HostOption {
	return func(h *HostOptions) {
		h.Target = target
	}
}

// WithHostOptionPorts speccifies what port will be cracked
func WithHostOptionPorts(p uint16) HostOption {
	return func(h *HostOptions) {
		h.Port = p
	}
}

// WithServiceOptions applies ServiceOptions to reffering host
// see service_options.go
func WithServiceOptions(opts ...ServiceOption) HostOption {
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
