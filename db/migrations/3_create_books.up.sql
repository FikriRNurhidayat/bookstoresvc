-- Adding new table called books
CREATE TABLE IF NOT EXISTS books (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  synopsis TEXT NOT NULL,
  isbn VARCHAR(255) NOT NULL,
  page_count INT NOT NULL,
  cover_type VARCHAR(100) NOT NULL,
  cover_image_url TEXT NOT NULL,
  author_id BIGINT NOT NULL REFERENCES authors(id),
  publisher_id BIGINT NOT NULL REFERENCES publishers(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  published_at TIMESTAMP NOT NULL
);
