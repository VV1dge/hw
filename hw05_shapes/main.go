package main

import (
	"errors"
	"fmt"
	"math"
)

// интерфейс Shape с методом для площади

type Shape interface {
	Area() float64
}

// структура для фигуры круг

type Circle struct {
	radius float64
}

// метод для вычисления площади круга

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// структура для фигуры прямоугольник

type Rectangle struct {
	width, height float64
}

// метод для вычисления площади прямоугольника

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// структура для фигуры треугольник

type Triangle struct {
	base, height float64
}

// метод для вычисления площади треугольника

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// функция, которая ожидает на входе объект типа Shape и возвращает его площадь

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("объект не реализует интерфейс Shape")
	}
	return shape.Area(), nil
}

func main() {
	c := Circle{5}
	r := Rectangle{5, 10}
	t := Triangle{8, 6}

	// вызов функции calculateArea для кажой фигуры

	areac, err := calculateArea(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Площадь фигуры", areac)
	arear, err := calculateArea(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Площадь фигуры", arear)
	areat, err := calculateArea(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Площадь фигуры", areat)
}
