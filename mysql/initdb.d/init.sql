DROP SCHEMA IF EXISTS ddd_menta;
CREATE SCHEMA ddd_menta;
USE ddd_menta;

DROP TABLE IF EXISTS users;

CREATE TABLE users
(
  user_id varchar(255) not null,
  name varchar(50) not null,
  email varchar(50) not null,
  password varchar(255) not null,
  profile text not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (user_id, name, email, password, profile) VALUES ("Menta", "test_user", "test@co.jp", "AJRUsjquq", "テストユーザーです");

DROP TABLE IF EXISTS user_careers;

CREATE TABLE user_careers
(
  user_career_id varchar(255) not null,
  user_id varchar(255) not null,
  `from` date not null,
  `to` date not null,
  detail text not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_careers (user_career_id, user_id, `from`, `to`, detail) VALUES ("Menta", "Menta", "2020-04-01", "2020-10-01", "ddd");

-- CREATE TABLE users
-- (
--   id int unsigned primary key auto_increment,
--   name varchar(50),
--   password varchar(255)
-- );

-- 実行されない場合
-- docker exec -it menta_db /bin/bash
-- mysql -uddd_menta -pddd_menta
-- use ddd_menta
-- create,insert実行
