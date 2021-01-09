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
    title TINYTEXT,
    content LONGTEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE board (
    id INT PRIMARY KEY,
    user_id VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE token (
    user_id VARCHAR(20) PRIMARY KEY, 
    refresh_token VARCHAR(32),
    refresh_token_expired_at TIMESTAMP,
    fcm_token VARCHAR(32),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

ALTER TABLE article add foreign key (article_user_id) references user(id);
ALTER TABLE article add foreign key (article_board_id) references board(id);
ALTER TABLE board add foreign key (board_user_id) references user(id);
ALTER TABLE token add foreign key (token_user_id) references user(id);
