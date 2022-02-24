package entity

import "time"

type Book struct {
  Id            uint32    `db:"id"`
  Title         string    `db:"title"`
  Synopsis      string    `db:"synopsis"`
  Isbn          string    `db:"isbn"`
  PageCount     uint32    `db:"page_count"`
  CoverType     string    `db:"cover_type"`
  CoverImageUrl string    `db:"cover_image_url"`
  AuthorId      uint32    `db:"author_id"`
  AuthorName    string    `db:"author_name"`
  PublisherId   uint32    `db:"publisher_id"`
  PublisherName string    `db:"publisher_name"`
  PublishedAt   time.Time `db:"published_at"`
}
