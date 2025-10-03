package main

import "fmt"
type Book struct {
 Title  string
 Author string
 Taken  bool
}

type Library struct {
 Books []Book
}

func (l *Library) AddBook(b Book) {
 l.Books = append(l.Books, b)
}

func (l *Library) FindByAuthor(author string) []Book {
 var res []Book
 for _, b := range l.Books {
  if b.Author == author {
   res = append(res, b)
  }
 }
 return res
}

func (l *Library) Borrow(title string) bool {
 for i := range l.Books {
  if l.Books[i].Title == title && !l.Books[i].Taken {
   l.Books[i].Taken = true
   return true
  }
 }
 return false
}

func (l *Library) Return(title string) {
 for i := range l.Books {
  if l.Books[i].Title == title {
   l.Books[i].Taken = false
  }
 }
}

func main() {
 lib := Library{}
 lib.AddBook(Book{"GO для чайников", "Галчанин В.С.", false})
 lib.Borrow("Go язык")
 fmt.Println(lib.Books)
}
