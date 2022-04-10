CREATE TABLE books (
                       id SERIAL NOT NULL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       publisher VARCHAR(255) NOT NULL,
                       authors VARCHAR[] NOT NULL
);

CREATE TABLE photos (
                        id SERIAL NOT NULL PRIMARY KEY,
                        photo TEXT NOT NULL,
                        book_id INT NOT NULL,
                        create_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        CONSTRAINT books
                            FOREIGN KEY(book_id)
                                REFERENCES photos(id)
);