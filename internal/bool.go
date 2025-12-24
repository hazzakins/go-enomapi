package internal

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// Bool decodes empty or non-standard boolean values without failing XML parsing.
type Bool bool

func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}
	parsed, err := parseBool(raw)
	if err != nil {
		return err
	}
	*b = Bool(parsed)
	return nil
}

func (b *Bool) UnmarshalText(text []byte) error {
	parsed, err := parseBool(string(text))
	if err != nil {
		return err
	}
	*b = Bool(parsed)
	return nil
}

func parseBool(value string) (bool, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return false, nil
	}
	switch strings.ToLower(value) {
	case "yes", "y":
		return true, nil
	case "no", "n":
		return false, nil
	}
	return strconv.ParseBool(value)
}
