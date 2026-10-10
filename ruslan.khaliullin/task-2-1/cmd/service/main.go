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

func main() {
}
