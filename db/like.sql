create database infinite_like;
use infinite_like;

#  把点赞者和文章对象关系存唯一索引

create table `like_record`(
     `id` bigint(20) unsigned not null auto_increment comment '主键ID',
     `biz_id` varchar(64)  not null default '' comment '业务ID',
     `obj_id` bigint(20) unsigned not null default 0 comment '点赞目标ID',
     `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
     `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型 0:点赞 1:点踩',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
     PRIMARY KEY (`id`),
     KEY `ix_update_time`( `update_time`),
     UNIQUE KEY `uk_biz_obj_uid` (`biz_id`,`obj_id`,`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='点赞记录表';

create table `like_count`(
    `id` bigint(20) unsigned not null auto_increment comment '主键ID',
    `biz_id` varchar(64)  not null default '' comment '业务ID',
    `obj_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '内容id',
    `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `dislike_num` int(11) NOT NULL DEFAULT '0' COMMENT '点踩数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time`( `update_time`),
    UNIQUE KEY `uk_biz_obj` (`biz_id`,`obj_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='点赞计数表';