CREATE TABLE `order` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `fid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '电影ID',
    `amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '订单金额',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态，0表示未支付，1表示已支付，2表示订单失效，默认为0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_fid` (`fid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
