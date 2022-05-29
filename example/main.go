package main

import (
	"context"
	"fmt"

	"github.com/myjupyter/ncrack"
)

func main() {
	c := ncrack.New(
		ncrack.WithPerHostOption(ncrack.HostOption{
			Protocols: []ncrack.ProtocolType{
				ncrack.ProtocolTypePSQL,
			},
			Target: "0.0.0.0",
			Ports:  []uint16{5432},
		}),
		ncrack.WithUser("postgres"),
		ncrack.WithPass("postgres"),
	)

	fmt.Println(c.Args())

	ctx := context.Background()
	buf, err := c.Run(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
