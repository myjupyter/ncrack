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
	args []string
	ctx  context.Context
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

// RunBytes synchronously executes ncrack,
func (c Cracker) RunBytes() ([]byte, error) {
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
	data, err := c.RunBytes()
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
	return res, nil
}

func (c Cracker) returnCtx() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}
