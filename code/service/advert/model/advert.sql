CREATE TABLE `advert`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255)  NOT NULL DEFAULT '' COMMENT '标题',
    `content` varchar(255)  NOT NULL DEFAULT '' COMMENT '内容',
    `is_com` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '广告性质，0表示普通公告，1表示商业广告，默认为0',
    `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '广告状态，0表示未失效，1表示已失效，默认为0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
