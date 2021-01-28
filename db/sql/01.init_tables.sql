-- CREATE DATABASE IF NOT EXISTS `webpracticedb`;
USE webpracticedb;

CREATE TABLE user (
    user_id VARCHAR(20) PRIMARY KEY,
    user_email VARCHAR(30) UNIQUE,
    user_pw VARCHAR(20),
    user_name VARCHAR(20),
    user_role VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE article (
    article_id INT PRIMARY KEY,
    board_id INT,
    user_id VARCHAR(20),
    article_title TINYTEXT,
    article_content LONGTEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE board (
    board_id INT PRIMARY KEY,
    board_name VARCHAR(20),
    user_id VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE token (
    user_id VARCHAR(20) PRIMARY KEY, 
    refresh_token VARCHAR(32) UNIQUE,
    refresh_token_expired_at TIMESTAMP,
    fcm_token VARCHAR(32) UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

ALTER TABLE article add foreign key (user_id) references user(user_id);
ALTER TABLE article add foreign key (board_id) references board(board_id);
ALTER TABLE board add foreign key (user_id) references user(user_id);
ALTER TABLE token add foreign key (user_id) references user(user_id);
