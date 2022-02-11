DROP SCHEMA IF EXISTS ddd_menta;
CREATE SCHEMA ddd_menta;
USE ddd_menta;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
  user_id VARCHAR(255) NOT NULL,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,
  profile TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id)
);

INSERT INTO users (
  user_id,
  name,
  email,
  password,
  profile
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "test_user",
  "test@co.jp",
  "AJRUsjquq",
  "テストユーザーです"
);

DROP TABLE IF EXISTS user_careers;

CREATE TABLE user_careers
(
  user_career_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  `from` DATETIME NOT NULL,
  `to` DATETIME NOT NULL,
  detail TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_career_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO user_careers (
  user_career_id,
  user_id,
  `from`,
  `to`,
  detail
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "2006-01-02 15:04:05",
  "2006-01-02 15:04:05",
  "ddd"
);

DROP TABLE IF EXISTS user_skills;
CREATE TABLE user_skills
(
  user_skill_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  tag VARCHAR(255) NOT NULL,
  assessment INT NOT NULL,
  experience_years INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_skill_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
INSERT INTO user_skills (
  user_skill_id,
  user_id,
  tag,
  assessment,
  experience_years
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "Go",
  1,
  1
);

DROP TABLE IF EXISTS mentors;
CREATE TABLE mentors
(
  mentor_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  mainImg VARCHAR(255) NOT NULL,
  subImg VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  detail TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (mentor_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO mentors (
  mentor_id,
  user_id,
  title,
  mainImg,
  subImg,
  category,
  detail
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "プログラミング全般のメンタリング",
  "/main.jpg",
  "/sub.jpg",
  "プログライミング",
  "設計・開発・テストの一覧をサポートできます"
);

DROP TABLE IF EXISTS mentor_skills;
CREATE TABLE mentor_skills
(
  mentor_skill_id VARCHAR(255) NOT NULL,
  mentor_id VARCHAR(255) NOT NULL,
  tag VARCHAR(255) NOT NULL,
  assessment INT NOT NULL,
  experience_years INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (mentor_skill_id),
  FOREIGN KEY (mentor_id) REFERENCES mentors(mentor_id)
);
INSERT INTO mentor_skills (
  mentor_skill_id,
  mentor_id,
  tag,
  assessment,
  experience_years
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "Go",
  5,
  5
);

DROP TABLE IF EXISTS plans;
CREATE TABLE plans
(
  plan_id VARCHAR(255) NOT NULL,
  mentor_id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  detail TEXT NOT NULL,
  plan_type INT NOT NULL,
  price INT NOT NULL,
  planStatus INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (plan_id),
  FOREIGN KEY (mentor_id) REFERENCES mentors(mentor_id)
);

INSERT INTO plans (
  plan_id,
  mentor_id,
  title,
  category,
  detail,
  plan_type,
  price,
  planStatus
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
	"DDDのメンタリング",
	"設計",
	"DDDの設計手法を学べます",
  2,
  1000,
  1
);

-- CREATE TABLE users
-- (
--   id INT unsigned PRIMARY KEY auto_increment,
--   name VARCHAR(50),
--   password VARCHAR(255)
-- );

-- 実行されない場合
-- docker exec -it menta_db /bin/bash
-- mysql -uddd_menta -pddd_menta
-- use ddd_menta
-- create,insert実行
-- select * from user_careers;
