DROP SCHEMA IF EXISTS ddd_menta;
CREATE SCHEMA ddd_menta;
USE ddd_menta;

DROP TABLE IF EXISTS users;

CREATE TABLE users
(
  user_id varchar(255) unsigned,
  name varchar(50),
  email varchar(50),
  password varchar(255),
  profile varchar(255),
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

INSERT INTO users (user_id, name, email, password, profile) VALUES ("Menta", "test_user", "test@co.jp", "AJRUsjquq", "テストユーザーです");

-- CREATE TABLE users
-- (
--   id int unsigned primary key auto_increment,
--   name varchar(50),
--   password varchar(255)
-- );
