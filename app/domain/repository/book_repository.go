package repository;

import (
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"

  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

type IBookRepository interface {
  GetBooks() ([]entity.Book, error)
  GetBook(id uint32) (entity.Book, error)
}

type BookRepository struct {
  db *sqlx.DB;
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (n BookRepository) GetBooks() ([]entity.Book, error) {
  var books []entity.Book

  statement, err := n.db.Prepare(
    "SELECT books.id, books.title, books.synopsis, books.isbn, books.page_count, books.cover_type, books.cover_image_url, " +
    "books.author_id, authors.name, " +
    "books.publisher_id, publishers.name, " +
    "books.published_at " + 
    "FROM books " +
    "JOIN authors ON authors.id = books.author_id " +
    "JOIN publishers ON publishers.id = books.publisher_id " +
    "ORDER BY books.id;")

  if err != nil {
    return nil, err
  }

  defer statement.Close()

  rows, err := statement.Query()
  if err != nil {
    return nil, err
  }

  for rows.Next() {
    book := entity.Book{}

    err = rows.Scan(
      &book.Id,
      &book.Title,
      &book.Synopsis,
      &book.Isbn,
      &book.PageCount,
      &book.CoverType,
      &book.CoverImageUrl,
      &book.AuthorId,
      &book.AuthorName,
      &book.PublisherId,
      &book.PublisherName,
      &book.PublishedAt,
    )
		if err != nil {
			return nil, err
		}

		books = append(books, book)
  }

  return books, nil
}

func (n BookRepository) GetBook(id uint32) (entity.Book, error) {
  var book entity.Book

  err := n.db.Get(
    &book,
    "SELECT books.id, books.title, books.synopsis, books.isbn, books.page_count, books.cover_type, books.cover_image_url, " +
    "books.author_id, authors.name AS author_name, " +
    "books.publisher_id, publishers.name AS publisher_name, " +
    "books.published_at " + 
    "FROM books " +
    "JOIN authors ON authors.id = books.author_id " +
    "JOIN publishers ON publishers.id = books.publisher_id " +
    "WHERE books.id = $1;",
    id,
  )

  if err != nil {
    return book, err
  }

  return book, nil
}
