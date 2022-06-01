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
			ncrack.WithHostOptionPorts(5437),
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypePostgres),
			ncrack.WithHostOptionPorts(5432),
			ncrack.WithHostOptionPorts(5433),
			ncrack.WithHostOptionTarget("scanner-dev.echelon.lan"),
			ncrack.WithHostOptionServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithHostOption(
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypeSSH),
			ncrack.WithHostOptionPorts(22),
			ncrack.WithHostOptionTarget("scanner-dev.echelon.lan"),
			ncrack.WithHostOptionServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("psql,ssh", "ftp"),
			ncrack.WithModuleOptionServiceOptions(
				ncrack.WithServiceOptionMinConnLimit(5),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("ssh"),
			ncrack.WithModuleOptionServiceOptions(
				ncrack.WithServiceOptionMinConnLimit(5),
			),
		),
		ncrack.WithUser("postgres", "admin", "psql", "echelon", "seclab"),
		ncrack.WithPass("postgres", "admin", "psql", "echelon", "seclab"),
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
