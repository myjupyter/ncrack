package main

import (
	"encoding/json"

	"fmt"

	"github.com/myjupyter/ncrack"
)

func main() {
	c := ncrack.New(
		ncrack.WithHostOption(
			ncrack.WithHostOptionPorts(5437),
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypePostgres.String()),
			ncrack.WithHostOptionPorts(5432),
			ncrack.WithHostOptionPorts(5433),
			ncrack.WithHostOptionTarget("localhost"),
			ncrack.WithServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithHostOption(
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypeSSH.String()),
			ncrack.WithHostOptionPorts(2222),
			ncrack.WithHostOptionTarget("localhost"),
			ncrack.WithServiceOptions(
				ncrack.WithServiceOptionConnDelay(1),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("psql,ssh", "ftp"),
			ncrack.WithModuleServiceOptions(
				ncrack.WithServiceOptionMinConnLimit(5),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("ssh"),
			ncrack.WithModuleServiceOptions(
				ncrack.WithServiceOptionMinConnLimit(5),
			),
		),
		ncrack.WithUser("postgres", "admin", "psql", "postgres"),
		ncrack.WithPass("postgres", "admin", "psql", "postgres"),
		ncrack.WithXMLOutput(),
	)

	fmt.Println(c.Args())

	res, err := c.Run()
	if err != nil {
		panic(err)
	}

	buf, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

}
