package ncrack

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
)

type Cracker struct {
	args []string
}

func New(opts ...Option) *Cracker {
	c := &Cracker{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c Cracker) Args() []string {
	return c.args
}

func (c Cracker) Run(ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "ncrack", c.args...)
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
