package main

import "fmt"

type Publisher struct {
	Name     string
	Founders []string
}

func (p Publisher) Totalfounders() int {
	return len(p.Founders)
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Publisher
}

func (b Book) showInfo() string {
	return fmt.Sprintf("%s was written by %s with %d pages and has %d Founders", b.Title, b.Author, b.Pages, b.Totalfounders())
}

type Movie struct {
	Title     string
	Director  string
	IMDBRting float64
	Publisher
}

func main() {

	// Create a new book

	book := Book{
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
		Pages:  224,
		Publisher: Publisher{
			Name:     "Pan Books",
			Founders: []string{"Cyril M. Kornbluth", "Robert A. Heinlein", "Isaac Asimov"},
		},
	}

	fmt.Println("Number of founders: ", book.Totalfounders())
	fmt.Println("Book info: ", book.showInfo())
}
