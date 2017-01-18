package client

import "encoding/xml"

type ResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`

	Header *ResponseHeader `xml:"Header"`
	Body   *ResponseBody   `xml:"Body"`
}

type ResponseHeader struct {
	Security *ResponseSecurity `xml:"Security"`
}

type ResponseSecurity struct {
	Username string `xml:"UsernameToken>Username"`
	Password string `xml:"UsernameToken>Password"`
}

type ResponseBody struct {
	GetAssets *ResponseGetAssets `xml:"Get_Assets_Response"`
	Error     *ResponseError     `xml:"Fault"`
}

type ResponseObjectList []*ResponseObject

type ResponseObject struct {
	Descriptor string              `xml:"Descriptor,attr,omitempty"`
	IDs        []*ResponseObjectID `xml:"ID"`
}

type ResponseObjectID struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type ResponseError struct {
	Code    string `xml:"faultcode"`
	Message string `xml:"faultstring"`
}

func (ol *ResponseObjectList) Contains(t, v string) bool {
	var l ResponseObjectList
	l = *ol

	for _, o := range l {
		for _, id := range o.IDs {
			if id.Type == t && id.Value == v {
				return true
			}
		}
	}
	return false
}

func (o *ResponseObject) Contains(t, v string) bool {
	for _, id := range o.IDs {
		if id.Type == t && id.Value == v {
			return true
		}
	}
	return false
}
