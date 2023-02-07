CREATE TABLE `comment` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `fid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '电影ID',
    `title` varchar(255)  NOT NULL DEFAULT '' COMMENT '标题',
    `content` varchar(255)  NOT NULL DEFAULT '' COMMENT '内容',
    `is_anonymous` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '匿名状态，0表示不匿名，1表示匿名，默认为0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_fid` (`fid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
