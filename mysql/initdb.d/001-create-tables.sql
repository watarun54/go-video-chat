set @now = NOW();

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

insert into users(name, email, password, created_at, updated_at) values("admin", "admin@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test1", "test1@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test2", "test2@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test3", "test3@test.com", "password", @now, @now);
insert into users(name, email, password, created_at, updated_at) values("test4", "test4@test.com", "password", @now, @now);


---- drop ----
DROP TABLE IF EXISTS `rooms`;

---- create ----
create table IF not exists `rooms` (
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into rooms(name, created_at, updated_at) values("room1", @now, @now);
insert into rooms(name, created_at, updated_at) values("room2", @now, @now);


---- drop ----
DROP TABLE IF EXISTS `user_rooms`;

---- create ----
create table IF not exists `user_rooms` (
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `room_id`          INT(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into user_rooms(user_id, room_id, created_at, updated_at) values(1, 1, @now, @now);
insert into user_rooms(user_id, room_id, created_at, updated_at) values(2, 1, @now, @now);
insert into user_rooms(user_id, room_id, created_at, updated_at) values(3, 1, @now, @now);
insert into user_rooms(user_id, room_id, created_at, updated_at) values(4, 2, @now, @now);
insert into user_rooms(user_id, room_id, created_at, updated_at) values(5, 2, @now, @now);


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

insert into comments(text, user_id, created_at, updated_at) values("comment1", 1, @now, @now);
insert into comments(text, user_id, created_at, updated_at) values("comment2", 1, @now, @now);
insert into comments(text, user_id, created_at, updated_at) values("comment3", 2, @now, @now);
insert into comments(text, user_id, created_at, updated_at) values("comment4", 2, @now, @now);
