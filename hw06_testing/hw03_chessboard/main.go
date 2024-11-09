package main

import "fmt"

func CreateBoard(x, y int) string {
	for i := 1; i <= x; i++ { // проходимся циклом по строкам
		for j := 1; j <= y; j++ { // проходмся циклом по столбцам
			if i%2 != 0 && j%2 != 0 || i%2 == 0 && j%2 == 0 {
				fmt.Print(" ") // так как пробел это белое, то по условию белые клетки
				// в четных строках на четных местах, а в нечетных на нечетных местах
			} else {
				fmt.Print("#") // в любом другом случае на оставшихся местах будут черные клетки, то есть #
			}
		}
		fmt.Print("\n\n")
	}
	return ""
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
