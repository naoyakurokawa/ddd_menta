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
  main_img VARCHAR(255) NOT NULL,
  sub_img VARCHAR(255) NOT NULL,
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
  main_img,
  sub_img,
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
  tag VARCHAR(255) NOT NULL,
  detail TEXT NOT NULL,
  plan_type INT NOT NULL,
  price INT NOT NULL,
  plan_status INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (plan_id),
  FOREIGN KEY (mentor_id) REFERENCES mentors(mentor_id)
);

INSERT INTO plans (
  plan_id,
  mentor_id,
  title,
  category,
  tag,
  detail,
  plan_type,
  price,
  plan_status
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
	"DDDのメンタリング",
	"設計",
	"DDD",
	"DDDの設計手法を学べます",
  2,
  1000,
  1
);

DROP TABLE IF EXISTS contracts;
CREATE TABLE contracts
(
  contract_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  mentor_id VARCHAR(255) NOT NULL,
  plan_id VARCHAR(255) NOT NULL,
  contract_status INT NOT NULL,
  created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (contract_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (plan_id) REFERENCES plans(plan_id)
);

INSERT INTO contracts (
  contract_id,
  user_id,
  mentor_id,
  plan_id,
  contract_status,
  created_at,
  updated_at
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad",
  1,
  "2022-02-22 00:00:00.000000",
  "2022-02-22 00:00:00.000000"
);

DROP TABLE IF EXISTS recruits;
CREATE TABLE recruits
(
  recruit_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  budget INT NOT NULL,
  recruit_type INT NOT NULL,
  detail TEXT NOT NULL,
  recruit_status INT NOT NULL,
  created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (recruit_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO recruits (
  recruit_id,
  user_id,
  title,
  budget,
  recruit_type,
  detail,
  recruit_status,
  created_at,
  updated_at
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "DDDの基礎を教えて下さい",
  5000,
  1,
  "DDDによる開発をサポートしてもらいたく募集しました",
  1,
  "2022-03-22 00:00:00.000000",
  "2022-03-22 00:00:00.000000"
);

DROP TABLE IF EXISTS suggestions;
CREATE TABLE suggestions
(
  suggestion_id VARCHAR(255) NOT NULL,
  mentor_id VARCHAR(255) NOT NULL,
  recruit_id VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  suggestion_type INT NOT NULL,
  detail TEXT NOT NULL,
  suggestion_status INT NOT NULL,
  created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (suggestion_id),
  FOREIGN KEY (mentor_id) REFERENCES mentors(mentor_id),
  FOREIGN KEY (recruit_id) REFERENCES recruits(recruit_id)
);

INSERT INTO suggestions (
  suggestion_id,
  mentor_id,
  recruit_id,
  price,
  suggestion_type,
  detail,
  suggestion_status,
  created_at,
  updated_at
) VALUES (
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  "e2e908dc-5981-4c4a-8e98-4487d3e122ad", 
  5000,
  1,
  "DDDの設計から開発までサポートします",
  1,
  "2022-03-22 00:00:00.000000",
  "2022-03-22 00:00:00.000000"
);
