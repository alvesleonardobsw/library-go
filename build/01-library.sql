CREATE DATABASE library;

\c library;

CREATE TABLE Author (
    IdAuthor VARCHAR(36) NOT NULL,
    Name VARCHAR(20) NOT NULL,
    LastName VARCHAR(20) NOT NULL,
    CONSTRAINT PK_Author PRIMARY KEY (IdAuthor)
);

CREATE TABLE Book (
    IdBook VARCHAR(36) NOT NULL,
    Name VARCHAR(20) NOT NULL,
    Genre VARCHAR(20) NOT NULL,
    IdAuthor VARCHAR(36) NOT NULL,
    CONSTRAINT PK_Book PRIMARY KEY (IdBook),
    CONSTRAINT FK_Book_Author FOREIGN KEY (IdAuthor) REFERENCES Author (IdAuthor)
);