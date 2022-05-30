package main

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/myjupyter/ncrack"
)

func main() {
	c := ncrack.New(
		ncrack.WithHostOption(
			ncrack.WithHostOptionProtocols(
				ncrack.ProtocolTypePSQL,
			),
			ncrack.WithHostOptionPorts(5432),
			ncrack.WithHostOptionTarget("0.0.0.0"),
			ncrack.WithHostOptionServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithHostOption(
			ncrack.WithHostOptionProtocols(
				ncrack.ProtocolTypePSQL,
			),
			ncrack.WithHostOptionPorts(5432),
			ncrack.WithHostOptionTarget("scanner-dev.echelon.lan"),
			ncrack.WithHostOptionServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithUser("postgres", "admin", "psql"),
		ncrack.WithPass("postgres", "admin", "psql"),
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
