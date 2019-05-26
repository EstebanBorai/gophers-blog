USE songs-share;

CREATE TABLE users (
  id varchar(36) PRIMARY KEY,
  firstName varchar(255),
  lastName varchar(255),
  username varchar(25),
  password varchar(25)
);

INSERT INTO users(id, firstName, lastName, username, password)
  VALUES('13067ca2-40bf-443e-b3fc-7281c011964f', 'John', 'Doe', 'johndoe', '12345678'),
        ('79f8a6aa-fbf0-4be1-b203-2ef71f172a43', 'Esteban', 'Borai', 'estebanborai', '12345678'),
        ('319ddde4-117a-49e4-82f2-ea7f3fc2b093', 'John', 'Snow', 'johnsnow', '12345678'),
        ('519dede3-112v-4e3t-98d2-ia7jri388u89', 'Another', 'User', 'anotheruser', '12345678')
