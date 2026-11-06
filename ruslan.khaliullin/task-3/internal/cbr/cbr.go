package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func Load(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	var curs valCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("parse input file: %w", err)
	}

	rates := make([]Currency, 0, len(curs.Valutes))
	for _, valuteItem := range curs.Valutes {
		if valuteItem.NumCode == "" || valuteItem.CharCode == "" {
			continue
		}

		numCode, err := strconv.Atoi(valuteItem.NumCode)
		if err != nil {
			return nil, fmt.Errorf("parse num code of %s: %w", valuteItem.CharCode, err)
		}

		value, err := strconv.ParseFloat(strings.Replace(valuteItem.Value, ",", ".", 1), 64)
		if err != nil {
			return nil, fmt.Errorf("parse value of %s: %w", valuteItem.CharCode, err)
		}

		rates = append(rates, Currency{
			NumCode:  numCode,
			CharCode: valuteItem.CharCode,
			Value:    value,
		})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}
