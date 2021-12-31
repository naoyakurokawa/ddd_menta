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
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  primary key (user_id)
);

INSERT INTO users (user_id, name, email, password, profile) VALUES ("e2e908dc-5981-4c4a-8e98-4487d3e122ad", "test_user", "test@co.jp", "AJRUsjquq", "テストユーザーです");

DROP TABLE IF EXISTS user_careers;

CREATE TABLE user_careers
(
  user_career_id varchar(255) not null,
  user_id varchar(255) not null,
  `from` datetime not null,
  `to` datetime not null,
  detail text not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  primary key (user_career_id),
  foreign key (user_id) references users(user_id)
);

INSERT INTO user_careers (user_career_id, user_id, `from`, `to`, detail) VALUES ("e2e908dc-5981-4c4a-8e98-4487d3e122ad", "e2e908dc-5981-4c4a-8e98-4487d3e122ad", "2006-01-02 15:04:05", "2006-01-02 15:04:05", "ddd");

DROP TABLE IF EXISTS user_skills;
CREATE TABLE user_skills
(
  user_skill_id varchar(255) not null,
  user_id varchar(255) not null,
  tag varchar(255) not null,
  assessment int not null,
  experience_years int not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  primary key (user_skill_id),
  foreign key (user_id) references users(user_id)
);
INSERT INTO user_skills (user_skill_id, user_id, tag, assessment, experience_years) VALUES ("e2e908dc-5981-4c4a-8e98-4487d3e122ad", "e2e908dc-5981-4c4a-8e98-4487d3e122ad", "Go", 1, 1);

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
-- select * from user_careers;
