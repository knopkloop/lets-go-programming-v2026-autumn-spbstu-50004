package main

import "fmt"

const (
	limitMin = 15
	limitMax = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() conditioner {
	return conditioner{low: limitMin, high: limitMax}
}

func (c *conditioner) regulate(sign string, value int) int {
	switch sign {
	case ">=":
		if value > c.low {
			c.low = value
		}
	case "<=":
		if value < c.high {
			c.high = value
		}
	}

	if c.low > c.high {
		return -1
	}

	return c.low
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for d := 0; d < departments; d++ {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		cond := newConditioner()

		for i := 0; i < employees; i++ {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				return
			}

			fmt.Println(cond.regulate(sign, value))
		}
	}
}
