CREATE TABLE `pay` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `oid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单ID',
    `amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '支付金额',
    `source` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式，0表示线下支付，1表示线上支付，默认为线下支付',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态，0表示未支付，1表示已支付，2表示支付失效，默认为0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_oid` (`oid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
