\c library;

-- Insert Authors

INSERT INTO Author(IdAuthor, Name, LastName)
VALUES ('1', 'Leo', 'Alves');

INSERT INTO Author(IdAuthor, Name, LastName)
VALUES ('2', 'Alef', 'Martins');

-- Insert Books

INSERT INTO Book (IdBook, Name, Genre, IdAuthor)
VALUES ('1', 'The walking Dead', 'Zumbie', '1');

INSERT INTO Book (IdBook, Name, Genre, IdAuthor)
VALUES ('2', 'Game Of Thrones', 'dragão', '1');

INSERT INTO Book (IdBook, Name, Genre, IdAuthor)
VALUES ('3', 'Demon Slayer', 'manga', '2');