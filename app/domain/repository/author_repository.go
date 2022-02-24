package repository

import (
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"

  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

type IAuthorRepository interface {
  GetAuthors() ([]entity.Author, error)
  GetAuthor(id uint32) (entity.Author, error)
}

type AuthorRepository struct {
  db *sqlx.DB;
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
  return &AuthorRepository{
    db: db,
  }
}

func (n AuthorRepository) GetAuthors() ([]entity.Author, error) {
  var authors []entity.Author

  statement, err := n.db.Prepare(
    "SELECT authors.id, authors.name " +
    "FROM authors " +
    "ORDER BY authors.id;")

  if err != nil {
    return nil, err
  }

  defer statement.Close()

  rows, err := statement.Query()
  if err != nil {
    return nil, err
  }

  for rows.Next() {
    author := entity.Author{}

    err = rows.Scan(
      &author.Id,
      &author.Name,
    )
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
  }

  return authors, nil
}

func (n AuthorRepository) GetAuthor(id uint32) (entity.Author, error) {
  var author entity.Author

  err := n.db.Get(
    &author,
    "SELECT authors.id, authors.name " +
    "FROM authors " +
    "ORDER BY authors.id;")
  if err != nil {
    return author, err
  }

  return author, nil
}
