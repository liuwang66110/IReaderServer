-- 用户注册
CREATE TABLE `user_main` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` varchar(64)  NOT NULL default '' COMMENT '名称',
    `password` varchar(60) not null default '' comment '',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态标识。1-可用，2-禁用, 3-删除',
    `token` varchar(40) not null default '' comment '访问token',
    `expired_at` timestamp not null default '2000-01-01 00:00:00' comment '登陆过期时间',
    `created_at` timestamp NOT NULL default '2000-01-01 00:00:00' COMMENT 'create time',
    `last_login_at`  timestamp NOT NULL default '2000-01-01 00:00:00'  COMMENT '最后登录时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`),
    UNIQUE KEY `uk_token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- 用户信息
CREATE TABLE `user_info` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint(20) unsigned not null default 0 comment 'user_main.id',
    `name` varchar(64)  NOT NULL default '' COMMENT '昵称',
    `mobile` varchar(64) NOT NULL DEFAULT '' COMMENT 'binding mobile',
    `gender` int unsigned not null default 1 comment '1-男; 2-女',
    `email` varchar(255) not null default '',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_usd` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- 好友关系
create table `user_friend` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint(20) unsigned not null default 0 comment 'user_main.id',
    `friend_id` bigint(20) unsigned not null default 0 comment 'user_main.id',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '1' comment '好友状态标识。1-已申请，2-已拒绝，3-好友',
    `content` varchar(255) not null default '' comment '内容',
    `confirm_at` timestamp NOT NULL default '2000-01-01 00:00:00'  COMMENT 'confirm time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
    UNIQUE KEY `uk_fid` (`user_id`, `friend_id`),
)ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;