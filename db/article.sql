create database infinite_article;
use infinite_article;

create table article(
    `id` bigint(20) unsigned not null auto_increment comment '主键ID',
    `title` varchar(255) not null default '' comment '标题',
    `content` text collate utf8_unicode_ci not null comment '内容',
    `cover` varchar(255) not null default '' comment '封面',
    `description` varchar(255) not null default '' comment '描述',
    `author_id` bigint(20) unsigned not null default 0 comment '作者ID',
    `status` tinyint(4) not null default 0 comment '状态 0:待审核 1:审核不通过 2:可见',
    `comment_num` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
    `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `collect_num` int(11) NOT NULL DEFAULT '0' COMMENT '收藏数',
    `view_num` int(11) NOT NULL DEFAULT '0' COMMENT '浏览数',
    `share_num` int(11) NOT NULL DEFAULT '0' COMMENT '分享数',
    `tag_ids`   varchar(255) not null default '' comment '标签ID',
    `publish_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_author_id` (`author_id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8mb4_bin  COMMENT='文章表';


