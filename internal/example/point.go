package main

import "fmt"

type Books struct {
	Page int
	Title string
}
func _() {
	books := Books{
		Page:  1,
		Title: "a",
	}
	var booksPoint *Books
	booksPoint = &books
	booksPoint.Page = 2
	fmt.Println(books)
}


