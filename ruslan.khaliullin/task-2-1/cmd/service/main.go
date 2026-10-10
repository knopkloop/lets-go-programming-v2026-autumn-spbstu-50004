package main

import "fmt"

const (
	defMinTemperature = 15
	defMaxTemperature = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() conditioner {
	return conditioner{low: defMinTemperature, high: defMaxTemperature}
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
		fmt.Println("Invalid department number")
		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Invalid employee number")
			return
		}

		cond := newConditioner()

		for range employees {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("Invalid temperature constraint")
				return
			}

			fmt.Println(cond.regulate(sign, value))
		}
	}
}
