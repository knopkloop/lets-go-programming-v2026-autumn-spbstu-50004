package main

import "fmt"

func main() {
	var n1, n2 int
	var op string

	_, err := fmt.Scan(&n1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&n2)
	if err != nil {
		fmt.Println("Ivalid second operand")
		return
	}

	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
}
