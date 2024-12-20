//Import packages for program
package main	
//Import fmt for I/O Operations
import "fmt"

//Design of this program is to create a library and be able to add various features
type Book struct {
	Title string
	Year int
	Author string
}

//Create a library with an empty slice of Book
type Library struct {
	Books []Book
}

//AddBook Method with a pointer receiver
func (l *Library) AddBook(book Book){
	l.Books = append(l.Books, book)
	
}

//ViewBooks Method
func (l Library) ViewBooks(){
	for _, book := range l.Books{
	fmt.Printf("Title: %s, Year %d, Author: %s\n",book.Title,book.Year,book.Author)
	} 
}

/*Create main func
-instantiate a library
-add books
-viewbook
*/
func main(){
	//instantiate library
	BigLibrary := Library{[]Book{}}
	//addbooks
	BigLibrary.AddBook(Book{Title: "The Gaming Guide", Year: 2022, Author: "Nicholas Alvarez"})
	//addbooks
	BigLibrary.AddBook(Book{Title: "The Youtube Creation Guide", Year: 2023, Author: "Nicholas Alvarez"})
	//viewbooks
	BigLibrary.ViewBooks()
}

// go run practice.go