package main

import "fmt"

func CreateBoard(x, y int) string {
	var chess string
	for i := 1; i <= x; i++ { // проходимся циклом по строкам
		for j := 1; j <= y; j++ { // проходмся циклом по столбцам
			if i%2 != 0 && j%2 != 0 || i%2 == 0 && j%2 == 0 {
				// так как пробел это белое, то по условию белые клетки
				// в четных строках на четных местах, а в нечетных на нечетных местах
				chess += " "
			} else {
				// в любом другом случае на оставшихся местах будут черные клетки, то есть #
				chess += "#"
			}
		}

		chess += "\n"
	}
	return chess
}

func main() {
	var x, y int
	fmt.Println("Введите значение по горизонтали")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println("Введите значение по вертикали")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Print(CreateBoard(x, y))
}