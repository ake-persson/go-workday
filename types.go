package client

import (
	"encoding/xml"
	"time"
)

type Boolean bool

type Date struct {
	time.Time
}

type DateTime struct {
	time.Time
}

func (b *Boolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var i int
	if bool(*b) {
		i = 1
	} else {
		i = 0
	}

	e.EncodeElement(i, start)

	return nil
}

func (d *Date) UnmarshalXML(decode *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decode.DecodeElement(&s, &start); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02-15:04", s)
	if err != nil {
		return err
	}
	*d = Date{t}

	return nil
}

func (d *DateTime) UnmarshalXML(decode *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decode.DecodeElement(&s, &start); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02-15:04", s)
	if err != nil {
		return err
	}
	*d = DateTime{t}

	return nil
}
