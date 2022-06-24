package ncrack

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"os/exec"
	"strings"
)

type Cracker struct {
	args       []string
	ctx        context.Context
	filterFunc func(*Service) bool
}

func New(opts ...Option) *Cracker {
	c := &Cracker{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Args method outputs arguments passed to ncrack
func (c Cracker) Args() []string {
	return c.args
}

// RunBytes synchronously executes ncrack and get
func (c Cracker) runBytes() ([]byte, error) {
	cmd := exec.CommandContext(c.returnCtx(), "ncrack", c.args...)
	buf, errBuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = errBuf
	if err := cmd.Run(); err != nil {
		if errBuf.Len() != 0 {
			return nil, errors.New(strings.TrimSpace(errBuf.String()))
		}
		return nil, err
	}
	return buf.Bytes(), nil
}

// Run synchronously executes ncrack
func (c Cracker) Run() (res *Run, err error) {
	data, err := c.runBytes()
	if err != nil {
		return nil, err
	}
	res = &Run{
		rawXML: data,
	}
	err = xml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	res.Services = c.filterServices(res.Services)
	return res, nil
}

func (c Cracker) returnCtx() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

func (c Cracker) filterServices(ss []Service) []Service {
	if c.filterFunc == nil {
		return ss
	}
	nss := []Service{}
	for i := range ss {
		if c.filterFunc(&ss[i]) {
			nss = append(nss, ss[i])
		}
	}
	return nss
}
