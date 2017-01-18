package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	XMLNsEnv  = "http://schemas.xmlsoap.org/soap/envelope/"
	XMLNsWsse = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	XMLNsWd   = "urn:com.workday/bsvc"
)

type RequestEnvelope struct {
	XMLName  xml.Name `xml:"env:Envelope"`
	XMLNsEnv string   `xml:"xmlns:env,attr"`

	Header *RequestHeader `xml:"env:Header"`
	Body   *RequestBody   `xml:"env:Body"`
}

type RequestHeader struct {
	Security *RequestSecurity `xml:"wsse:Security"`
}

type RequestSecurity struct {
	XMLNsWsse string `xml:"xmlns:wsse,attr"`

	Username string `xml:"wsse:UsernameToken>wsse:Username"`
	Password string `xml:"wsse:UsernameToken>wsse:Password"`
}

type RequestBody struct {
	Request *Request
}

type Request struct {
	*RequestGetAssets
}

type RequestObjectList []*RequestObject

type RequestObject struct {
	Descriptor string             `xml:"wd:Descriptor,attr,omitempty"`
	IDs        []*RequestObjectID `xml:"wd:ID"`
}

type RequestObjectID struct {
	Type  string `xml:"wd:type,attr"`
	Value string `xml:",chardata"`
}

func NewObjectList() *RequestObjectList {
	return &RequestObjectList{}
}

func (rl *RequestObjectList) Add(descr string, ids []*RequestObjectID) {
	r := RequestObject{
		Descriptor: descr,
		IDs:        ids,
	}

	*rl = append(*rl, &r)
}

func (c *Client) newEnvelope() *RequestEnvelope {
	return &RequestEnvelope{
		XMLNsEnv: XMLNsEnv,
		Header: &RequestHeader{
			Security: &RequestSecurity{
				XMLNsWsse: XMLNsWsse,
				Username:  c.Username + "@" + c.Tenant,
				Password:  c.Password,
			},
		},
		Body: &RequestBody{},
	}
}

func (c *Client) getRequestXML(req *Request) ([]byte, error) {
	env := c.newEnvelope()
	env.Body.Request = req

	buf, err := xml.MarshalIndent(env, "", "  ")
	if err != nil {
		return []byte{}, nil
	}

	return buf, nil
}

func (c *Client) getResponse(reqBody *Request) (*http.Response, error) {
	var url = c.APIURL + "/" + c.Tenant + "/Resource_Management/" + c.APIVersion

	reqEnv := c.newEnvelope()
	reqEnv.Body.Request = reqBody

	b := &bytes.Buffer{}
	encoder := xml.NewEncoder(b)
	if err := encoder.Encode(reqEnv); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) get(reqBody *Request) (*ResponseEnvelope, error) {
	resp, err := c.getResponse(reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var env ResponseEnvelope
	if err := xml.NewDecoder(resp.Body).Decode(&env); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %d message: %s", resp.StatusCode, env.Body.Error.Message)
	}

	return &env, nil
}

func (c *Client) getXML(reqBody *Request) ([]byte, error) {
	resp, err := c.getResponse(reqBody)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var env ResponseEnvelope
		if err := xml.NewDecoder(resp.Body).Decode(&env); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("request failed with status: %d message: %s", resp.StatusCode, env.Body.Error.Message)
	}

	return ioutil.ReadAll(resp.Body)
}
