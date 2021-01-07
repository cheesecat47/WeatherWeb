-- CREATE DATABASE IF NOT EXISTS `webpracticedb`;
USE webpracticedb;

CREATE TABLE user (
    id VARCHAR(20) PRIMARY KEY,
    pw VARCHAR(20),
    name VARCHAR(20),
    role VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE article (
    id INT PRIMARY KEY,
    board_id INT,
    user_id VARCHAR(20),
    title VARCHAR(50),
    content VARCHAR(200),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE board (
    id INT PRIMARY KEY,
    user_id VARCHAR(20)
);

CREATE TABLE token (
    id VARCHAR(20),
    refresh_token VARCHAR(20),
    refresh_token_expired_at TIMESTAMP,
    fcm_token VARCHAR(20)
);

ALTER TABLE article add foreign key (user_id) references user(id);
ALTER TABLE article add foreign key (board_id) references board(id);
ALTER TABLE board add foreign key (user_id) references user(id);

INSERT INTO user values ('jy', 'pw', 'sjy', null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
insert INTO user values ('tg', 'pw2', 'atg', null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);