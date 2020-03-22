CREATE USER  'gomicro'@'%'  IDENTIFIED BY  'gomicro';
grant all privileges on gomicro.* to gomicro@'%' identified by 'gomicro'; flush privileges;

CREATE DATABASE `gomicro`;
USE `gomicro`;

CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20)   NOT NULL COMMENT '用户名',
    `pwd`          varchar(128)  NOT NULL COMMENT '密码',
    `created_time` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
    `updated_time` TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB COMMENT ='用户表';

INSERT INTO user (user_id, user_name, pwd) VALUE (10001, 'micro', '123');
