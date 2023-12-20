create database infinite_user;
use infinite_user;

create table user(
    `id` bigint(20) UNSIGNED not null auto_increment COMMENT '主键id',
    `username` varchar(32) not null default '' COMMENT '用户名',
    `avatar` varchar(256) not null default '' COMMENT '头像',
    `mobile` varchar(128) NOT NULL default '' comment '手机号',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    primary key(id),
    KEY `ix_update_time` (`update_time`),
    UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';