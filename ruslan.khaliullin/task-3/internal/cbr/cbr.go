package cbr

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	if err := xml.Unmarshal(data, &curs); err != nil {
		return nil, fmt.Errorf("parse input file: %w", err)
	}

	rates := make([]Currency, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		numCode, err := strconv.Atoi(v.NumCode)
		if err != nil {
			return nil, fmt.Errorf("parse num code of %s: %w", v.CharCode, err)
		}

		value, err := strconv.ParseFloat(strings.Replace(v.Value, ",", ".", 1), 64)
		if err != nil {
			return nil, fmt.Errorf("parse value of %s: %w", v.CharCode, err)
		}

		if math.IsNaN(value) || math.IsInf(value, 0) {
			return nil, fmt.Errorf("value of %s is not a finite number", v.CharCode)
		}

		rates = append(rates, Currency{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}
