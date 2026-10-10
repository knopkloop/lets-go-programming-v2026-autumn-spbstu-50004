package main

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
}
