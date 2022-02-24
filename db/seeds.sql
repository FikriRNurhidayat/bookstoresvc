-- Seeding Publishers
INSERT INTO publishers (id, name) VALUES
  (1, 'Toha Putra'),
  (2, 'Balai Pustaka');

-- Seeding Authors
INSERT INTO authors (id, name)
VALUES
  (1, 'Karl Marx'),
  (2, 'Frederich Nietzsche');

-- Seeding Books
INSERT INTO books (id, title, synopsis, isbn, page_count, cover_type, cover_image_url, author_id, publisher_id, published_at)
VALUES 
  (
    1,
    'Industrial Society and its future',
    'Lorem ipsum dolor sit amet',
    '12123123123',
    100,
    'PAPERBACK',
    'https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/old-books-arranged-on-shelf-royalty-free-image-1572384534.jpg',
    1,
    1,
    NOW()
  ),
  (
    2,
    'Industrial Society Or Something',
    'Lorem ipsum dolor sit amet',
    '12123123123',
    100,
    'PAPERBACK',
    'https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/old-books-arranged-on-shelf-royalty-free-image-1572384534.jpg',
    1,
    2,
    NOW()
  ),
  (
    3,
    'Industrial Society',
    'Lorem ipsum dolor sit amet',
    '12123123123',
    100,
    'PAPERBACK',
    'https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/old-books-arranged-on-shelf-royalty-free-image-1572384534.jpg',
    2,
    2,
    NOW()
  ),
  (
    4,
    'Future',
    'Lorem ipsum dolor sit amet',
    '12123123123',
    100,
    'PAPERBACK',
    'https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/old-books-arranged-on-shelf-royalty-free-image-1572384534.jpg',
    2,
    1,
    NOW()
  );


-- Reset authors_id_seq
SELECT setval('authors_id_seq', (COALESCE((SELECT authors.id FROM authors ORDER BY authors.id DESC LIMIT 1), 0) + 1), true);

-- Reset publishers_id_seq
SELECT setval('publishers_id_seq', (COALESCE((SELECT publishers.id FROM publishers ORDER BY publishers.id DESC LIMIT 1), 0) + 1), true);

-- Reset books_id_seq
SELECT setval('books_id_seq', (COALESCE((SELECT books.id FROM books ORDER BY books.id DESC LIMIT 1), 0) + 1), true);
