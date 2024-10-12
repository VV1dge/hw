package main

import "fmt"

// Book структура c полями.
type Book struct {
	id     int32
	title  string
	author string
	year   int32
	size   int32
	rate   float32
}

// Select выбор режима сравнения.
type Select int

const (
	ByYear Select = iota
	BySize
	ByRate
)

// BookSravn структура  с полями для сравнения.
type BookSravn struct {
	Year int32
	Size int32
	Rate float32
	mode Select
}

func NewBookSravn(year int32, size int32, rate float32, mode Select) *BookSravn {
	return &BookSravn{
		Year: year,
		Size: size,
		Rate: rate,
		mode: mode,
	}
}
func (b *BookSravn) Compare(other *BookSravn) bool {
	switch b.mode {
	case ByYear:
		return b.Year > other.Year
	case BySize:
		return b.Size > other.Size
	case ByRate:
		return b.Rate > other.Rate
	default:
		return false
	}
}

// GetID метод для получения поля.
func (r *Book) GetID() int32 {
	return r.id
}

// SetID метод для установки поля.
func (r *Book) SetID(id int32) {
	r.id = id
}
func (r *Book) GetTitle() string {
	return r.title
}
func (r *Book) SetTitle(title string) {
	r.title = title
}
func (r *Book) GetAuthor() string {
	return r.author
}
func (r *Book) SetAuthor(author string) {
	r.author = author
}
func (r *Book) GetYear() int32 {
	return r.year
}
func (r *Book) SetYear(year int32) {
	r.year = year
}
func (r *Book) GetSize() int32 {
	return r.size
}
func (r *Book) SetSize(size int32) {
	r.size = size
}
func (r *Book) GetRate() float32 {
	return r.rate
}
func (r *Book) SetRate(rate float32) {
	r.rate = rate
}

func main() {
	// book1 создание экземпляра структуры Book
	book1 := Book{}
	book1.SetID(1)
	book1.SetTitle("Lukomore")
	book1.SetAuthor("Pushkin")
	book1.SetYear(1830)
	book1.SetSize(10)
	book1.SetRate(1.5)
	fmt.Println("id:", book1.GetID(), "title:", book1.GetTitle(), "author:", book1.GetAuthor(),
		"year:", book1.GetYear(), "size:", book1.GetSize(), "rate:", book1.GetRate())
	booksr1 := NewBookSravn(1830, 2, 1.1, ByRate)
	booksr2 := NewBookSravn(1831, 1, 1.0, ByRate)
	fmt.Println("Сравнение по рейтингу:", booksr1.Compare(booksr2))
	booksr3 := NewBookSravn(1830, 3, 1.0, BySize)
	booksr4 := NewBookSravn(1831, 5, 1.0, BySize)
	fmt.Println("Сравнение по размеру:", booksr3.Compare(booksr4))
	booksr5 := NewBookSravn(1831, 1, 1.0, ByYear)
	booksr6 := NewBookSravn(1830, 4, 1.0, ByYear)
	fmt.Println("Сравнение по году:", booksr5.Compare(booksr6))
}
