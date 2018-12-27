-- 用户注册
CREATE TABLE `user_main` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'admin id',
    `name` varchar(64)  NOT NULL default '' COMMENT '名称',
    `mobile` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'binding mobile',
    `password` varchar(60) not null default '' comment '',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态标识。1-可用，2-禁用, 3-删除',
    `token` varchar(40) not null default '' comment '访问token',
    `expired_at` timestamp not null default '2000-01-01 00:00:00' comment '登陆过期时间',
    `created_at` timestamp NOT NULL default '2000-01-01 00:00:00' COMMENT 'create time',
    `last_login_at`  timestamp NOT NULL default '2000-01-01 00:00:00'  COMMENT '最后登录时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_mobile` (`mobile`),
    UNIQUE KEY `uk_name` (`name`),
    UNIQUE KEY `uk_token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;