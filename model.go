package ncrack

import (
	"encoding/xml"
	"strconv"
	"time"
)

type Run struct {
	XMLName xml.Name `xml:"ncrackrun" json:"-"`

	Args             string    `xml:"args,attr" json:"args"`
	Start            Timestamp `xml:"start,attr" json:"start"`
	StartStr         string    `xml:"startstr,attr" json:"startstr"`
	Version          string    `xml:"version,attr" json:"version"`
	XMLOutputVersion string    `xml:"xmloutputversion,attr" json:"xmloutputversion"`

	Services []Service `xml:"service" json:"service"`

	rawXML []byte `xml:"-" json:"-"`
}

func (r Run) Bytes() []byte {
	return r.rawXML
}

type Service struct {
	StartTime Timestamp `xml:"starttime,attr" json:"starttime"`
	EndTime   Timestamp `xml:"endtime,attr" json:"endtime"`

	Address     Address     `xml:"address" json:"address"`
	Port        Port        `xml:"port" json:"port"`
	Credentials Credentials `xml:"credentials" json:"credentials"`
}

type Address struct {
	Addr string `xml:"addr,attr" json:"addr"`
	Type string `xml:"addrtype,attr" json:"addrtype"`
}

type Port struct {
	Protocol string `xml:"protocol,attr" json:"protocol"`
	Port     uint16 `xml:"portid,attr" json:"portid"`
	Name     string `xml:"name,attr" json:"name"`
}

type Credentials struct {
	Username string `xml:"username,attr" json:"username"`
	Password string `xml:"password,attr" json:"password"`
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
