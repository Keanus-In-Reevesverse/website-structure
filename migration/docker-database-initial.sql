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
game_image VARCHAR(200) NOT NULL,
video VARCHAR(60) NOT NULL,
PRIMARY KEY (name, publisher),
FOREIGN KEY (genre_ID) REFERENCES GENRE(genre_ID)
);
-- ALTER TABLE GAME ADD FOREIGN KEY(genre_ID) REFERENCES GENRE (genre_ID)

CREATE TABLE USER (
user_ID INT auto_increment PRIMARY KEY,
email VARCHAR(60) UNIQUE,
name VARCHAR(60),
phone_number VARCHAR(12),
password VARCHAR(100)
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
change_ID int auto_increment PRIMARY KEY,
game_ID INT,
store_name VARCHAR(60),
price DECIMAL(6,2),
change_date VARCHAR(60),
FOREIGN KEY(game_ID) REFERENCES GAME(game_ID),
FOREIGN KEY (store_name) REFERENCES STORE(name)
);
-- FOREIGN KEY(/*erro: ??*/) REFERENCES STORE (store_ID,name)