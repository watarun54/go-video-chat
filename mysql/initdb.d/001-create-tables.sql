---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users` (
 `id`               INT(20) AUTO_INCREMENT,
 `first_name`       VARCHAR(20) NOT NULL,
 `last_name`        VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

set @now = NOW();

insert into users(first_name, last_name, created_at, updated_at) values("Patricia", "Smith", @now, @now);
insert into users(first_name, last_name, created_at, updated_at) values("Linda", "Johnson", @now, @now);
insert into users(first_name, last_name, created_at, updated_at) values("Mary", "William", @now, @now);
insert into users(first_name, last_name, created_at, updated_at) values("Robert", "Jones", @now, @now);
insert into users(first_name, last_name, created_at, updated_at) values("James", "Brown", @now, @now);
