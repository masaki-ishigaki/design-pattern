package iterator

/*****************************
 * Book
 ****************************/
type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{
		name,
	}
}

func (book *Book) String() string {
	return book.name
}

/*****************************
 * BookIterator
 ****************************/
type BookIterator struct {
	bookshelf *BookShelf
	index     int
}

func NewBookIterator(shelf *BookShelf) *BookIterator {
	return &BookIterator{
		bookshelf: shelf,
		index:     0,
	}
}

func (iter *BookIterator) HasNext() bool {
	if iter.index < iter.bookshelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (iter *BookIterator) Next() *Book {
	book := iter.bookshelf.GetBookAt(iter.index)
	iter.index++
	return book
}

/*****************************
 * BookShelf
 ****************************/
type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf(maxSize int) *BookShelf {
	return &BookShelf{
		books: make([]Book, maxSize),
		last:  0,
	}
}

func (shelf *BookShelf) GetBookAt(index int) *Book {
	return &shelf.books[index]
}

func (shelf *BookShelf) AppendBook(book *Book) {
	if shelf.last < len(shelf.books) {
		shelf.books[shelf.last] = *book
		shelf.last++
	}
}

func (shelf *BookShelf) GetLength() int {
	return len(shelf.books)
}

func (shelf *BookShelf) Iterator() *BookIterator {
	return NewBookIterator(shelf)
}
