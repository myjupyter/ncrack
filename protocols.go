package ncrack

//go:generate go run github.com/dmarkham/enumer -type=ProtocolType -trimprefix=ProtocolType -transform=lower
type ProtocolType int

const (
	ProtocolTypeUnspecified ProtocolType = iota
	ProtocolTypeSSH
	ProtocolTypePSQL
	ProtocolTypeRDP
	ProtocolTypeFTP
	ProtocolTypeTelnet
	ProtocolTypeHTTP
	ProtocolTypeHTTPS
	ProtocolTypeWordpress
	ProtocolTypePOP3
	ProtocolTypePOP3S
	ProtocolTypeIMAP
	ProtocolTypeCVS
	ProtocolTypeSMB
	ProtocolTypeVNC
	ProtocolTypeSIP
	ProtocolTypeRedis
	ProtocolTypeMQTT
	ProtocolTypeMySQL
	ProtocolTypeMSSQL
	ProtocolTypeMongoDB
	ProtocolTypeCassandra
	ProtocolTypeWinRM
	ProtocolTypeOWA
	ProtocolTypeDICO
)
