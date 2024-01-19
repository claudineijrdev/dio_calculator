package main

import "fmt"

func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func Sub(n ...int) int {
	sub := n[0]
	for _, v := range n[1:] {
		sub -= v
	}
	return sub
}

func Multiply(n ...int) int {
	mul := 1
	for _, v := range n {
		mul *= v
	}
	return mul
}

func Divide(n ...int) (int, error) {
	div := n[0]
	for _, v := range n[1:] {
		if v == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		div /= v
	}
	return div, nil
}

func main() {
	sumResult := Sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sumResult)

	subResult := Sub(90, 15, 3)
	fmt.Println("Sub:", subResult)

	mulResult := Multiply(9, 4)
	fmt.Println("Multiply:", mulResult)

	divResult, err := Divide(81, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Divide:", divResult)

}
