# ncrack

##### This library provides possibility to use `ncrack` through golang interface

---

# Why ncrack ?

Ncrack is an open source tool for network authentication cracking. It was designed for high-speed parallel cracking using a dynamic engine that can adapt to different network situations. Ncrack can also be extensively fine-tuned for special cases, though the default parameters are generic enough to cover almost every situation. It is built on a modular architecture that allows for easy extension to support additional protocols. Ncrack is designed for companies and security professionals to audit large networks for default or weak passwords in a rapid and reliable way. It can also be used to conduct fairly sophisticated and intensive brute force attacks against individual services.

 *[Ncrack Official website](https://nmap.org/ncrack/man.html)*

##### Ncrack supports next protocols and moudules
- SSH
- RDP
- FTP
- Telnet
- HTTP(S)
- POP3(S)
- IMAP
- SMB
- VNC
- SIP
- Redis
- PostgreSQL
- MySQL
- MSSQL
- MongoDB
- Cassandra 
- WinRM
- OWA

---

# TODO list

- [ ] Add all ncrack options
- [ ] Write tests
- [ ] Add asynchronous cracking with progress output 

---

# Example 

```go
package main

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/myjupyter/ncrack"
)

func main() {
	c := ncrack.New(
		ncrack.WithContext(context.Background()),
		ncrack.WithHostOption(
			ncrack.WithHostOptionPorts(5437),
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypePostgres.String()),
			ncrack.WithHostOptionPorts(5432),
			ncrack.WithHostOptionPorts(5433),
			ncrack.WithHostOptionTarget("localhost"),
			ncrack.WithServiceOptions(
				ncrack.WithServiceOptionConnDelay("1"),
			),
		),
		ncrack.WithHostOption(
			ncrack.WithHostOptionProtocols(ncrack.ProtocolTypeSSH.String()),
			ncrack.WithHostOptionPorts(2222),
			ncrack.WithHostOptionTarget("localhost"),
			ncrack.WithServiceOptions(
				ncrack.WithServiceOptionConnDelay("1"),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("psql"),
			ncrack.WithModuleServiceOptions(
				ncrack.WithServiceOptionMinConnLimit(5),
				ncrack.WithServiceOptionConnDelay("30"),
			),
		),
		ncrack.WithModuleOption(
			ncrack.WithModuleOptionProtocol("ssh"),
			ncrack.WithModuleServiceOptions(
				ncrack.WithServiceOptionAuthTries(5),
				ncrack.WithServiceOptionMinConnLimit(5),
			),
		),
		ncrack.WithUser("postgres", "admin", "psql", "postgres"),
		ncrack.WithPass("postgres", "admin", "psql", "postgres"),
		ncrack.WithXMLOutput(),
	)

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
```
