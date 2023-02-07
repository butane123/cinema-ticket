CREATE TABLE `film` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '电影名称',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '电影描述',
    `stock` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '影票库存',
    `amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '影票金额',
    `screenwriter` varchar(255)  NOT NULL DEFAULT '' COMMENT '影片编剧',
    `director` varchar(255)  NOT NULL DEFAULT '' COMMENT '影片导演',
    `length` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '影片时长',
    `is_select_seat` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否支持选座，1表示可以，0表示不可以',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
