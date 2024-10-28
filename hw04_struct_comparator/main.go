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

type Sravn struct {
	mode Select
}

func NewSravn(mode Select) *Sravn {
	return &Sravn{
		mode: mode,
	}
}

func (b *Sravn) Compare(book1, book2 *Book) bool {
	switch b.mode {
	case ByYear:
		return book1.year > book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
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
	booktest := Book{}
	booktest.SetID(1)
	booktest.SetTitle("Lukomore")
	booktest.SetAuthor("Pushkin")
	booktest.SetYear(1830)
	booktest.SetSize(10)
	booktest.SetRate(1.5)
	fmt.Println("id:", booktest.GetID(), "title:", booktest.GetTitle(), "author:", booktest.GetAuthor(),
		"year:", booktest.GetYear(), "size:", booktest.GetSize(), "rate:", booktest.GetRate())
	book1 := Book{id: 123, title: "Groza", author: "Ostrovskiy", year: 1859, size: 333, rate: 3.5}
	book2 := Book{id: 124, title: "Oblomov", author: "Goncharov", year: 1847, size: 222, rate: 3.6}
	booksravnByYear := NewSravn(ByYear)
	fmt.Println(booksravnByYear.Compare(&book1, &book2))
	booksravnBySize := NewSravn(BySize)
	fmt.Println(booksravnBySize.Compare(&book1, &book2))
	booksravnByRate := NewSravn(ByRate)
	fmt.Println(booksravnByRate.Compare(&book1, &book2))
}
