-- https://github.com/kukaro/Eris-DockerExampleTemplate/blob/master/Mysql.InitSql/mysql-init-files/create.sql
-- http://jameskim79.ipdisk.co.kr:8000/apps/wordpress/mariadb-%EB%AA%85%EB%A0%B9%EC%96%B4-%EB%B0%8F-%EC%82%AC%EC%9A%A9%EB%B2%95/

create database db default character set utf8;

CREATE TABLE db.accounts(
    id VARCHAR(20) PRIMARY KEY,
    pw VARCHAR(20),
    name VARCHAR (100),
    age INT
);
