CREATE TABLE `pro_user`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`     varchar(32)  DEFAULT NULL,
    `password`     varchar(128) DEFAULT NULL,
    `avatar`       varchar(64)  DEFAULT NULL,
    `role`         tinyint(1) DEFAULT '1' COMMENT '用户角色 1: admin 2: 开发',
    `status`       tinyint(1) DEFAULT '1' COMMENT '1: 启用 2: 禁用',
    `introduction` mediumtext COMMENT '介绍',
    `create_at`    datetime     DEFAULT CURRENT_TIMESTAMP,
    `update_at`    datetime     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_name` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;