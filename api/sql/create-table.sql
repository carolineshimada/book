CREATE DATABASE IF NOT EXISTS rest-api;
USE rest-api;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS fallower;
DROP TABLE IF EXISTS user;

CREATE TABLE user(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;


CREATE TABLE fallower(
    user_id int not null,
    fallower_id int not null,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (fallower_id) REFERENCES user(id) ON DELETE CASCADE,
    primary key(user_id, fallower_id)
) ENGINE=INNODB;


CREATE TABLE posts(
    id int auto_increment primary key,
    title varchar(50) not null,
    post varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id) REFERENCES user(id) ON DELETE CASCADE,
    likes int default 0,
    createdAt timestamp default current_timestamp
) ENGINE=InnoDB;

