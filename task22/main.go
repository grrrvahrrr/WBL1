package main

import (
	"fmt"
	"log"
)

const (
	Add   = "+" //adding numbers
	Sub   = "-" //subtracting numbers
	Multi = "*" //multiplying numbers
	Div   = "/" //squared numbers
)

//Structs
type Numbers struct {
	num1, num2 uint64
}

//Number Scanner
func Num2Scan() (a uint64, b uint64, err error) {
	fmt.Println(`Input two positive numbers, please (divided by space)`)
	_, err = fmt.Scanf("%d %d", &a, &b)
	return a, b, err
}

//Number Methods
func (n *Numbers) Addition() uint64 {
	return n.num1 + n.num2
}

func (n *Numbers) Subtraction() uint64 {
	return n.num1 - n.num2
}

func (n *Numbers) Multiplication() uint64 {
	return n.num1 * n.num2
}

func (n *Numbers) Division() uint64 {
	return n.num1 / n.num2
}

func main() {

	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "/" for divison`)

	var operation string
	var nums Numbers
	_, err := fmt.Scan(&operation)
	if err != nil {
		log.Fatal(err)
	}
	nums.num1, nums.num2, err = Num2Scan()
	if err != nil {
		log.Fatal(err)
	}

	switch operation {
	case Add:
		fmt.Println("Answer is: ", nums.Addition())
	case Sub:
		fmt.Println("Answer is: ", nums.Subtraction())
	case Multi:
		fmt.Println("Answer is: ", nums.Multiplication())
	case Div:
		fmt.Println("Answer is: ", nums.Division())

	default:
		log.Fatal(`Try again, something went wrong.. Invalid Operation`)
	}
}
