---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users` (
 `id`               BIGINT AUTO_INCREMENT,
 `name`             VARCHAR(255) NOT NULL,
 `email`            VARCHAR(255) NOT NULL,
 `hashed_password`  VARCHAR(255) NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into users(name, email, hashed_password) values("admin", "admin@test.com", "ca104e39a10ee19d78dcd1523b628c3028ba11343f355f81d4bea00065829161"); -- "password"
insert into users(name, email, hashed_password) values("test1", "test1@test.com", "0216c89866cd37f4356f39853add6e4abe120e997ac2613885647a024f700628"); -- "password"
insert into users(name, email, hashed_password) values("test2", "test2@test.com", "4f8f72d1f674af4804339a51499b09e3f377fb6c81ff0170b2c76737f03bef19"); -- "password"
insert into users(name, email, hashed_password) values("test3", "test3@test.com", "0a7cff99941bebb860307bc3785cd624a4e7033a883b2c7575ff76909c6f5972"); -- "password"
insert into users(name, email, hashed_password) values("test4", "test4@test.com", "16c634a74bc93c234fb3acf8f108de29beee1c12023cf56133237d85217fb21b"); -- "password"


---- drop ----
DROP TABLE IF EXISTS `rooms`;

---- create ----
create table IF not exists `rooms` (
 `id`               BIGINT AUTO_INCREMENT,
 `name`             VARCHAR(255) NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into rooms(name) values("room1");
insert into rooms(name) values("room2");


---- drop ----
DROP TABLE IF EXISTS `user_rooms`;

---- create ----
create table IF not exists `user_rooms` (
 `id`               BIGINT AUTO_INCREMENT,
 `user_id`          BIGINT NOT NULL,
 `room_id`          BIGINT NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into user_rooms(user_id, room_id) values(1, 1);
insert into user_rooms(user_id, room_id) values(2, 1);
insert into user_rooms(user_id, room_id) values(3, 1);
insert into user_rooms(user_id, room_id) values(4, 2);
insert into user_rooms(user_id, room_id) values(5, 2);


---- drop ----
DROP TABLE IF EXISTS `comments`;

---- create ----
create table IF not exists `comments` (
 `id`               BIGINT AUTO_INCREMENT,
 `text`             VARCHAR(255) NOT NULL,
 `user_id`          BIGINT NOT NULL,
 `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into comments(text, user_id) values("comment1", 1);
insert into comments(text, user_id) values("comment2", 1);
insert into comments(text, user_id) values("comment3", 2);
insert into comments(text, user_id) values("comment4", 2);
