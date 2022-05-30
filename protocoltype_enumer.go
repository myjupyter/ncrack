// Code generated by "enumer -type=ProtocolType -trimprefix=ProtocolType -transform=lower"; DO NOT EDIT.

package ncrack

import (
	"fmt"
	"strings"
)

const _ProtocolTypeName = "unspecifiedsshpsqlrdpftptelnethttphttpswordpresspop3pop3simapcvssmbvncsipredismqttmysqlmssqlmongodbcassandrawinrmowadico"

var _ProtocolTypeIndex = [...]uint8{0, 11, 14, 18, 21, 24, 30, 34, 39, 48, 52, 57, 61, 64, 67, 70, 73, 78, 82, 87, 92, 99, 108, 113, 116, 120}

const _ProtocolTypeLowerName = "unspecifiedsshpsqlrdpftptelnethttphttpswordpresspop3pop3simapcvssmbvncsipredismqttmysqlmssqlmongodbcassandrawinrmowadico"

func (i ProtocolType) String() string {
	if i < 0 || i >= ProtocolType(len(_ProtocolTypeIndex)-1) {
		return fmt.Sprintf("ProtocolType(%d)", i)
	}
	return _ProtocolTypeName[_ProtocolTypeIndex[i]:_ProtocolTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ProtocolTypeNoOp() {
	var x [1]struct{}
	_ = x[ProtocolTypeUnspecified-(0)]
	_ = x[ProtocolTypeSSH-(1)]
	_ = x[ProtocolTypePSQL-(2)]
	_ = x[ProtocolTypeRDP-(3)]
	_ = x[ProtocolTypeFTP-(4)]
	_ = x[ProtocolTypeTelnet-(5)]
	_ = x[ProtocolTypeHTTP-(6)]
	_ = x[ProtocolTypeHTTPS-(7)]
	_ = x[ProtocolTypeWordpress-(8)]
	_ = x[ProtocolTypePOP3-(9)]
	_ = x[ProtocolTypePOP3S-(10)]
	_ = x[ProtocolTypeIMAP-(11)]
	_ = x[ProtocolTypeCVS-(12)]
	_ = x[ProtocolTypeSMB-(13)]
	_ = x[ProtocolTypeVNC-(14)]
	_ = x[ProtocolTypeSIP-(15)]
	_ = x[ProtocolTypeRedis-(16)]
	_ = x[ProtocolTypeMQTT-(17)]
	_ = x[ProtocolTypeMySQL-(18)]
	_ = x[ProtocolTypeMSSQL-(19)]
	_ = x[ProtocolTypeMongoDB-(20)]
	_ = x[ProtocolTypeCassandra-(21)]
	_ = x[ProtocolTypeWinRM-(22)]
	_ = x[ProtocolTypeOWA-(23)]
	_ = x[ProtocolTypeDICO-(24)]
}

var _ProtocolTypeValues = []ProtocolType{ProtocolTypeUnspecified, ProtocolTypeSSH, ProtocolTypePSQL, ProtocolTypeRDP, ProtocolTypeFTP, ProtocolTypeTelnet, ProtocolTypeHTTP, ProtocolTypeHTTPS, ProtocolTypeWordpress, ProtocolTypePOP3, ProtocolTypePOP3S, ProtocolTypeIMAP, ProtocolTypeCVS, ProtocolTypeSMB, ProtocolTypeVNC, ProtocolTypeSIP, ProtocolTypeRedis, ProtocolTypeMQTT, ProtocolTypeMySQL, ProtocolTypeMSSQL, ProtocolTypeMongoDB, ProtocolTypeCassandra, ProtocolTypeWinRM, ProtocolTypeOWA, ProtocolTypeDICO}

var _ProtocolTypeNameToValueMap = map[string]ProtocolType{
	_ProtocolTypeName[0:11]:         ProtocolTypeUnspecified,
	_ProtocolTypeLowerName[0:11]:    ProtocolTypeUnspecified,
	_ProtocolTypeName[11:14]:        ProtocolTypeSSH,
	_ProtocolTypeLowerName[11:14]:   ProtocolTypeSSH,
	_ProtocolTypeName[14:18]:        ProtocolTypePSQL,
	_ProtocolTypeLowerName[14:18]:   ProtocolTypePSQL,
	_ProtocolTypeName[18:21]:        ProtocolTypeRDP,
	_ProtocolTypeLowerName[18:21]:   ProtocolTypeRDP,
	_ProtocolTypeName[21:24]:        ProtocolTypeFTP,
	_ProtocolTypeLowerName[21:24]:   ProtocolTypeFTP,
	_ProtocolTypeName[24:30]:        ProtocolTypeTelnet,
	_ProtocolTypeLowerName[24:30]:   ProtocolTypeTelnet,
	_ProtocolTypeName[30:34]:        ProtocolTypeHTTP,
	_ProtocolTypeLowerName[30:34]:   ProtocolTypeHTTP,
	_ProtocolTypeName[34:39]:        ProtocolTypeHTTPS,
	_ProtocolTypeLowerName[34:39]:   ProtocolTypeHTTPS,
	_ProtocolTypeName[39:48]:        ProtocolTypeWordpress,
	_ProtocolTypeLowerName[39:48]:   ProtocolTypeWordpress,
	_ProtocolTypeName[48:52]:        ProtocolTypePOP3,
	_ProtocolTypeLowerName[48:52]:   ProtocolTypePOP3,
	_ProtocolTypeName[52:57]:        ProtocolTypePOP3S,
	_ProtocolTypeLowerName[52:57]:   ProtocolTypePOP3S,
	_ProtocolTypeName[57:61]:        ProtocolTypeIMAP,
	_ProtocolTypeLowerName[57:61]:   ProtocolTypeIMAP,
	_ProtocolTypeName[61:64]:        ProtocolTypeCVS,
	_ProtocolTypeLowerName[61:64]:   ProtocolTypeCVS,
	_ProtocolTypeName[64:67]:        ProtocolTypeSMB,
	_ProtocolTypeLowerName[64:67]:   ProtocolTypeSMB,
	_ProtocolTypeName[67:70]:        ProtocolTypeVNC,
	_ProtocolTypeLowerName[67:70]:   ProtocolTypeVNC,
	_ProtocolTypeName[70:73]:        ProtocolTypeSIP,
	_ProtocolTypeLowerName[70:73]:   ProtocolTypeSIP,
	_ProtocolTypeName[73:78]:        ProtocolTypeRedis,
	_ProtocolTypeLowerName[73:78]:   ProtocolTypeRedis,
	_ProtocolTypeName[78:82]:        ProtocolTypeMQTT,
	_ProtocolTypeLowerName[78:82]:   ProtocolTypeMQTT,
	_ProtocolTypeName[82:87]:        ProtocolTypeMySQL,
	_ProtocolTypeLowerName[82:87]:   ProtocolTypeMySQL,
	_ProtocolTypeName[87:92]:        ProtocolTypeMSSQL,
	_ProtocolTypeLowerName[87:92]:   ProtocolTypeMSSQL,
	_ProtocolTypeName[92:99]:        ProtocolTypeMongoDB,
	_ProtocolTypeLowerName[92:99]:   ProtocolTypeMongoDB,
	_ProtocolTypeName[99:108]:       ProtocolTypeCassandra,
	_ProtocolTypeLowerName[99:108]:  ProtocolTypeCassandra,
	_ProtocolTypeName[108:113]:      ProtocolTypeWinRM,
	_ProtocolTypeLowerName[108:113]: ProtocolTypeWinRM,
	_ProtocolTypeName[113:116]:      ProtocolTypeOWA,
	_ProtocolTypeLowerName[113:116]: ProtocolTypeOWA,
	_ProtocolTypeName[116:120]:      ProtocolTypeDICO,
	_ProtocolTypeLowerName[116:120]: ProtocolTypeDICO,
}

var _ProtocolTypeNames = []string{
	_ProtocolTypeName[0:11],
	_ProtocolTypeName[11:14],
	_ProtocolTypeName[14:18],
	_ProtocolTypeName[18:21],
	_ProtocolTypeName[21:24],
	_ProtocolTypeName[24:30],
	_ProtocolTypeName[30:34],
	_ProtocolTypeName[34:39],
	_ProtocolTypeName[39:48],
	_ProtocolTypeName[48:52],
	_ProtocolTypeName[52:57],
	_ProtocolTypeName[57:61],
	_ProtocolTypeName[61:64],
	_ProtocolTypeName[64:67],
	_ProtocolTypeName[67:70],
	_ProtocolTypeName[70:73],
	_ProtocolTypeName[73:78],
	_ProtocolTypeName[78:82],
	_ProtocolTypeName[82:87],
	_ProtocolTypeName[87:92],
	_ProtocolTypeName[92:99],
	_ProtocolTypeName[99:108],
	_ProtocolTypeName[108:113],
	_ProtocolTypeName[113:116],
	_ProtocolTypeName[116:120],
}

// ProtocolTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ProtocolTypeString(s string) (ProtocolType, error) {
	if val, ok := _ProtocolTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ProtocolTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ProtocolType values", s)
}

// ProtocolTypeValues returns all values of the enum
func ProtocolTypeValues() []ProtocolType {
	return _ProtocolTypeValues
}

// ProtocolTypeStrings returns a slice of all String values of the enum
func ProtocolTypeStrings() []string {
	strs := make([]string, len(_ProtocolTypeNames))
	copy(strs, _ProtocolTypeNames)
	return strs
}

// IsAProtocolType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ProtocolType) IsAProtocolType() bool {
	for _, v := range _ProtocolTypeValues {
		if i == v {
			return true
		}
	}
	return false
}