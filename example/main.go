package main

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/myjupyter/ncrack"
)

func main() {
	c := ncrack.New(
		ncrack.WithPerHostOption([]ncrack.HostOption{
			{
				Protocols: []ncrack.ProtocolType{
					ncrack.ProtocolTypePSQL,
				},
				Target: "0.0.0.0",
				Ports:  []uint16{5432},
			},
			{
				Protocols: []ncrack.ProtocolType{
					ncrack.ProtocolTypePSQL,
				},
				Target: "scanner-dev.echelon.lan",
				Ports:  []uint16{5432},
			},
		}...),
		ncrack.WithUser("postgres"),
		ncrack.WithPass("postgres"),
		ncrack.WithXMLOutput(),
	)

	fmt.Println(c.Args())

	ctx := context.Background()
	buf, err := c.Run(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	r := &ncrack.Run{}
	err = xml.Unmarshal(buf, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v+", r)
}
