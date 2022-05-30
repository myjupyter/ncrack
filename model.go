package ncrack

import (
	"encoding/xml"
	"strconv"
	"time"
)

type Run struct {
	XMLName xml.Name `xml:"ncrackrun"`

	Args             string    `xml:"args,attr"`
	Start            Timestamp `xml:"start,attr"`
	StartStr         string    `xml:"startstr,attr"`
	Version          string    `xml:"version,attr"`
	XMLOutputVersion string    `xml:"xmloutputversion,attr"`

	Services []Service `xml:"service"`
}

type Service struct {
	StartTime Timestamp `xml:"starttime,attr"`
	EndTime   Timestamp `xml:"endtime,attr"`

	Address     Address     `xml:"address"`
	Port        Port        `xml:"port"`
	Credentials Credentials `xml:"credentials"`
}

type Address struct {
	Addr string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
}

type Port struct {
	Protocol string `xml:"protocol,attr"`
	Port     uint16 `xml:"portid,attr"`
	Name     string `xml:"psql,attr"`
}

type Credentials struct {
	Username string `xml:"username,attr"`
	Password string `xml:"password,attr"`
}

// Timestamp represents time as a UNIX timestamp in seconds.
type Timestamp time.Time

// ParseTime converts a UNIX timestamp string to a time.Time.
func (t *Timestamp) ParseTime(s string) error {
	timestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(timestamp, 0))

	return nil
}

// FormatTime formats the time.Time value as a UNIX timestamp string.
func (t Timestamp) FormatTime() string {
	return strconv.FormatInt(time.Time(t).Unix(), 10)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(t.FormatTime()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	return t.ParseTime(string(b))
}

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (t Timestamp) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if time.Time(t).IsZero() {
		return xml.Attr{}, nil
	}

	return xml.Attr{Name: name, Value: t.FormatTime()}, nil
}

// UnmarshalXMLAttr implements the xml.UnmarshalXMLAttr interface.
func (t *Timestamp) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	return t.ParseTime(attr.Value)
}
