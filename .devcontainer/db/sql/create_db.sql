CREATE database go;

use go;

CREATE TABLE users (
    ID VARCHAR(255) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL
);

INSERT INTO users (id, name, email) VALUES ("userid-1", "user1", "test1@example.com"), ("userid-2", "user2", "test2@example.com")