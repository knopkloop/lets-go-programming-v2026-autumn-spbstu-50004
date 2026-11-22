package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  xmlInt     `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    xmlFloat64 `json:"value"     xml:"Value"`
}

type xmlInt int

func (x *xmlInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode xml int: %w", err)
	}

	raw = strings.TrimSpace(raw)
	if raw == "" {
		*x = 0

		return nil
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("parse num code %q: %w", raw, err)
	}

	*x = xmlInt(val)

	return nil
}

type xmlFloat64 float64

func (x *xmlFloat64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode xml float: %w", err)
	}

	raw = strings.Replace(strings.TrimSpace(raw), ",", ".", 1)

	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", raw, err)
	}

	*x = xmlFloat64(val)

	return nil
}

type valCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

func Load(path string) ([]Currency, error) {
	safePath := filepath.Clean(path)

	data, err := os.ReadFile(safePath)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	var curs valCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("parse input file: %w", err)
	}

	rates := curs.Valutes

	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}
