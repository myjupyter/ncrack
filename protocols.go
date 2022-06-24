package ncrack

type ProtocolType string

func (pt ProtocolType) String() string {
	return string(pt)
}

const (
	ProtocolTypeUnspecified  ProtocolType = ""
	ProtocolTypeNetbiosSSN   ProtocolType = "netbios-ssn"
	ProtocolTypeMySQL        ProtocolType = "mysql"
	ProtocolTypeMicrosoftDS  ProtocolType = "microsoft-ds"
	ProtocolTypeMongoDB      ProtocolType = "mongodb"
	ProtocolTypePostgres     ProtocolType = "psql"
	ProtocolTypePOP3         ProtocolType = "pop3"
	ProtocolTypeSSH          ProtocolType = "ssh"
	ProtocolTypeMSSQL        ProtocolType = "mssql"
	ProtocolTypeOWA          ProtocolType = "owa"
	ProtocolTypeMSWBTServer  ProtocolType = "ms-wbt-server"
	ProtocolTypeWeb          ProtocolType = "web"
	ProtocolTypeTelnet       ProtocolType = "telnet"
	ProtocolTypeCassandra    ProtocolType = "cassandra"
	ProtocolTypeCVS          ProtocolType = "cvs"
	ProtocolTypeWP           ProtocolType = "wp"
	ProtocolTypeWebform      ProtocolType = "webform"
	ProtocolTypeMQTT         ProtocolType = "mqtt"
	ProtocolTypeJoomla       ProtocolType = "joomla"
	ProtocolTypeSMB2         ProtocolType = "smb2"
	ProtocolTypeWordPressTLS ProtocolType = "wordpress-tls"
	ProtocolTypePOP3S        ProtocolType = "pop3s"
	ProtocolTypeIMAP         ProtocolType = "imap"
	ProtocolTypeWinRM        ProtocolType = "winrm"
	ProtocolTypeWebTLS       ProtocolType = "web-tls"
	ProtocolTypeVNC          ProtocolType = "vnc"
	ProtocolTypeCouchbase    ProtocolType = "couchbase"
	ProtocolTypeHTTPS        ProtocolType = "https"
	ProtocolTypeWordPress    ProtocolType = "wordpress"
	ProtocolTypeHTTP         ProtocolType = "http"
	ProtocolTypeFTP          ProtocolType = "ftp"
	ProtocolTypeSIP          ProtocolType = "sip"
	ProtocolTypeRedis        ProtocolType = "redis"
	ProtocolTypeWPTLS        ProtocolType = "wp-tls"
	ProtocolTypeSMB          ProtocolType = "smb"
	ProtocolTypeDICOM        ProtocolType = "dicom"
	ProtocolTypeRPD          ProtocolType = "rdp"
)
