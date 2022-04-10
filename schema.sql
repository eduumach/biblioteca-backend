CREATE TABLE books (
                       ID SERIAL PRIMARY KEY,
                       Title VARCHAR NOT NULL,
                       Publisher VARCHAR NOT NULL,
                       Photo VARCHAR NOT NULL,
                       Authors VARCHAR[] NOT NULL
);

INSERT INTO books(title, publisher, photo, authors)
VALUES (
           'o amor',
           'editora teste',
           'foto.teste/o_amor',
           '{"joão","maria"}'
       );

SELECT * FROM books;