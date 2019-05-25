USE songs-share;

CREATE TABLE users (
  id BINARY(16) PRIMARY KEY,
  firstName varchar(255),
  lastName varchar(255),
  username varchar(25)
);

INSERT INTO users(id, firstName, lastName, username)
  VALUES(UUID_TO_BIN(UUID()), 'John', 'Doe', 'johndoe'),
        (UUID_TO_BIN(UUID()), 'Esteban', 'Borai', 'estebanborai'),
        (UUID_TO_BIN(UUID()), 'John', 'Snow', 'johnsnow')
