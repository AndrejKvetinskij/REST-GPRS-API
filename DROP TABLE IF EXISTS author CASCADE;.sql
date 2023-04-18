-- DROP TABLE IF EXISTS author CASCADE;
-- DROP TABLE IF EXISTS book CASCADE;
-- DROP TABLE IF EXISTS book_authors CASCADE;


-- CREATE TABLE author (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(100) NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
-- );
-- CREATE INDEX idx_author_create_at_pagination ON public.author (create_at, id);
-- CREATE INDEX idx_author_age_pagination ON public.author (age, id);


-- CREATE TABLE book (

--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(100) NOT NULL,
--     age INT,
--     is_alive BOOL,
--     created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')

-- );
-- CREATE INDEX idx_author_create_at_pagination ON public.book (create_at, id);
-- CREATE INDEX idx_author_age_pagination ON public.book (age, id);

-- CREATE TABLE book_authors (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     book_id UUID NOT NULL,
--     author_id UUID NOT NULL,

--     CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id),
--     CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id),
--     CONSTRAINT book_authors_unique UNIQUE (book_id, author_id)
-- );


-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Народ, 2022, true, 2023-04-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Джоан Роулинг, 30, true, 2023-04-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Джек Лондон, 30, true, 2023-05-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Джек, 45, false, 2023-05-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Джонс, 43, false, 2023-06-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ( Маин Рид, 41, true, 2023-06-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ( Макс Фрай, 39, false, 2023-07-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES(Антуан Дэ-Экзюпери, 37, true, 2023-08-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Александр Дюма, 35, false, 2023-08-09 16:28:33.737322);
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES (Иван Мележ, 33, true, 2023-08-09 16:28:33.737322);





-- DROP TABLE IF EXISTS author CASCADE;
-- DROP TABLE IF EXISTS book CASCADE;
-- DROP TABLE IF EXISTS book_authors CASCADE;
-- CREATE TABLE author (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(100) NOT NULL,
--     age INTEGER,
--     is_alive BOOLEAN,
--     created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
-- );
-- CREATE INDEX idx_author_created_at_pagination ON public.author (created_at, id);
-- CREATE INDEX idx_author_age_pagination ON public.author (age, id);
-- ;
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Народ', 2022, true, '2023-04-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Джоан Роулинг', 30, true, '2023-04-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Джек Лондон', 20, true, '2023-05-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Джек', 10, false, '2023-05-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Джонс', 40, false, '2023-06-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Маин Рид', 50, true, '2023-06-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Макс Фрай', 60, false, '2023-07-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES('АнтуанДэ-Экзюпери', 70, true, '2023-08-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Александр Дюма', 80, false, '2023-08-09 16:28:33.737322');
-- INSERT INTO author (name, age, is_alive, created_at)
-- VALUES ('Иван Мележ', 100, true, '2023-08-09 16:28:33.737322')


---SELECT COUNT(*)
FROM author;

OFFSETS
-- SELECT name, age, is_alive, created_at
-- FROM author
-- ORDER BY age DESC

-- ;

-- SELECT name,
--     age,
--     is_alive,
--     created_at
-- FROM author
-- ORDER BY age DESC
-- OFFSET 3
-- LIMIT 3;
