---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users` (
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `email`            VARCHAR(20) NOT NULL,
 `password`         VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

set @now = NOW();

insert into users(name, email, password, created_at, updated_at) values("admin", "admin@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test1", "test1@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test2", "test2@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test3", "test3@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test4", "test4@test.com", "password", @now, @now);


---- drop ----
DROP TABLE IF EXISTS `comments`;

---- create ----
create table IF not exists `comments` (
 `id`               INT(20) AUTO_INCREMENT,
 `text`             VARCHAR(20) NOT NULL,
 `user_id`          INT(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

set @now = NOW();

insert into comments(text, user_id, created_at, updated_at) values("comment1", 1, @now, @now);
insert into comments(text, user_id, created_at, updated_at) values("comment2", 1, @now, @now);
insert into comments(text, user_id, created_at, updated_at) values("comment3", 1, @now, @now);
