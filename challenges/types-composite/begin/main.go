// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct{
	name string
}

// define a book type with a title and author
type book struct{
	author
	title string
}
func (b book) getAuthorName() string {
	return b.author.name
}

// define a library type to track a list of books
type library struct {
	books map[string][]book
}

// define addBook to add a book to the library
func (l *library) addBook(b book) {
	l.books[b.author.name] = append(l.books[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l *library) lookupByAuthor(name string) []book {
	return l.books[name]
}

func main() {
	// create a new library
	l := library{
		books: make(map[string][]book),
	}

	// add 2 books to the library by the same author
	b1 := book{
		author: author{
			name: "James Bond",
		},
		title: "Foo",
	}

	b2 := book{
		author: author{
			name: "James Bond",
		},
	}
	l.addBook(b1)
	l.addBook(b2)

	// add 1 book to the library by a different author
	b3 := book {
		author: author{
			name: "Quixk",
		},
		title: "Sleep",
	}
	l.addBook(b3)

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	books := l.lookupByAuthor("James Bond")

	// print out the first book's title and its author's name
	fmt.Println(books[0])

}
