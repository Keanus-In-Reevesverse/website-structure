-- create table personalities(
--     id serial primary key,
--     name varchar,
--     story varchar
-- );

-- INSERT INTO personalities(name, story) VALUES
-- ('Eiichiro Oda', 'Eiichiro Oda is a Japanese manga artist and the creator of the series One Piece (1997–present). One Piece is both the best-selling manga in history and the best-selling comic series printed in volume.'),
-- ('Shigeru Miyamoto', 'Shigeru Miyamoto is a Japanese video game designer, producer and game director at Nintendo, where he serves as one of its representative directors.');

-- CREATE DATABASE BUSCAJOGOS;
-- DROP DATABASE BUSCAJOGOS;

-- USE BUSCAJOGOS;

CREATE TABLE GENRE (
genre_ID INT auto_increment PRIMARY KEY,
description VARCHAR(60) UNIQUE NOT NULL
);

CREATE TABLE GAME (
game_ID INT auto_increment UNIQUE,
genre_ID INT NOT NULL,
name VARCHAR(60) NOT NULL,
publisher VARCHAR(60) NOT NULL,
description VARCHAR(300),
game_image VARCHAR(200),
video VARCHAR(60) NOT NULL,
PRIMARY KEY (name, publisher),
FOREIGN KEY (genre_ID) REFERENCES GENRE(genre_ID)
);
-- ALTER TABLE GAME ADD FOREIGN KEY(genre_ID) REFERENCES GENRE (genre_ID)

CREATE TABLE USER (
user_ID INT auto_increment PRIMARY KEY,
email VARCHAR(60) UNIQUE,
name VARCHAR(60),
phone_number VARCHAR(12)
);

CREATE TABLE FAVORITE (
user_id INT,
game_id INT,
PRIMARY KEY (user_id, game_id),
FOREIGN KEY (user_id) REFERENCES USER(user_ID),
FOREIGN KEY (game_id) REFERENCES GAME(game_ID)
);
-- ALTER TABLE FAVORITE ADD FOREIGN KEY(user_id) REFERENCES USER (user_ID)
-- ALTER TABLE FAVORITE ADD FOREIGN KEY(game_id) REFERENCES GAME (game_ID)

CREATE TABLE LOGIN (
user_ID INT PRIMARY KEY,
email VARCHAR(60) UNIQUE,
password VARCHAR(30),
FOREIGN KEY (user_id) REFERENCES USER(user_ID)
);
-- FOREIGN KEY(email,,) REFERENCES USER (user_ID,email)

CREATE TABLE PRICE_ALERT (
user_ID INT,
game_ID INT,
price DECIMAL(6,2),
PRIMARY KEY (user_ID, game_ID),
FOREIGN KEY (user_id) REFERENCES USER(user_ID),
FOREIGN KEY (game_id) REFERENCES GAME(game_ID)
);
-- FOREIGN KEY(/*erro: ??*/) REFERENCES USER (user_ID,email)/*falha: chave estrangeira*/

CREATE TABLE STORE (
store_ID INT auto_increment PRIMARY KEY,
name VARCHAR(60) UNIQUE
);

CREATE TABLE GAME_PRICES (
game_ID INT,
store_name VARCHAR(60),
current_price DECIMAL(6,2),
PRIMARY KEY (game_ID, store_name),
FOREIGN KEY (store_name) REFERENCES STORE(name),
FOREIGN KEY (game_ID) REFERENCES GAME(game_ID)
);
-- ALTER TABLE GAME_PRICES ADD FOREIGN KEY(/*erro: ??*/) REFERENCES STORE (store_ID,name)


CREATE TABLE HISTORY (
game_ID INT,
store_name VARCHAR(60),
price DECIMAL(6,2),
change_date VARCHAR(60),
FOREIGN KEY(game_ID) REFERENCES GAME(game_ID),
FOREIGN KEY (store_name) REFERENCES STORE(name)
);
-- FOREIGN KEY(/*erro: ??*/) REFERENCES STORE (store_ID,name)


INSERT INTO GENRE VALUES (NULL, 'ACAO');

INSERT INTO GAME VALUES (NULL, (SELECT GENRE_ID FROM GENRE WHERE DESCRIPTION = 'ACAO'), 'FOR HONOR', 'UBSOFT', 'FOR HONOR É UM JOGO', NULL, 'LINK FDC');

INSERT INTO USER VALUES(NULL, 'EMAIL@FATEC', 'NOME ALEATORIO', '11945473106');

INSERT INTO FAVORITE VALUES((SELECT USER_ID FROM USER WHERE EMAIL = 'EMAIL@FATEC'), (SELECT GAME_ID FROM GAME WHERE NAME = 'FOR HONOR'));

INSERT INTO LOGIN VALUES ( (SELECT USER_ID FROM USER WHERE EMAIL = 'EMAIL@FATEC'), 'EMAIL@FATEC', 'Senha'); 

INSERT INTO PRICE_ALERT VALUES ( (SELECT USER_ID FROM USER WHERE EMAIL = 'EMAIL@FATEC'), (SELECT GAME_ID FROM GAME WHERE NAME = 'FOR HONOR'), 15.50);

INSERT INTO STORE VALUES (NULL, 'EPIC STORE');
INSERT INTO STORE VALUES (NULL, 'STEAM');

INSERT INTO GAME_PRICES VALUES ((SELECT GAME_ID FROM GAME WHERE NAME = 'FOR HONOR'), 'STEAM', 20.99);


INSERT INTO HISTORY VALUES ((SELECT GAME_ID FROM GAME WHERE NAME = 'FOR HONOR'), 'STEAM', 20.99, UNIX_TIMESTAMP());