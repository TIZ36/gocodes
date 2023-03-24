use gomock;

create table if not exists `basic_info`(
    `uid` bigint unsigned NOT NULL,
    `cur_stage` int unsigned NOT NULL DEFAULT 0 COMMENT '',
    `maincity_level` int unsigned NOT NULL DEFAULT 0 COMMENT '',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `user_info`(
    `uid` bigint unsigned NOT NULL,
    `data` blob NOT NULL COMMENT '',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `user_group_map` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `gid` bigint unsigned NOT NULL DEFAULT 0,
    `uid` bigint unsigned NOT NULL DEFAULT 0,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `user_friend` (
    `one_uid` bigint unsigned NOT NULL DEFAULT 0,
    `ano_uid` bigint unsigned NOT NULL DEFAULT 0,
    `relate_hash` char(32) NOT NULL DEFAULT 0,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`relate_hash`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `user_blacklist` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `one_uid` bigint unsigned NOT NULL DEFAULT 0,
    `ano_uid` bigint unsigned NOT NULL DEFAULT 0,
    `relate_hash` char(32) NOT NULL DEFAULT 0 COMMENT '快速找到两方的关系',
    `black_state` tinyint NOT NULL DEFAULT 0 COMMENT 'one_uid 主动方，0：未拉黑；1：拉黑',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE black_relation(`one_uid`, `ano_uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `group_owner` (
    `gid` bigint unsigned NOT NULL ,
    `owner_uid` bigint unsigned NOT NULL DEFAULT 0,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`gid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `group_config` (
    `group_id` bigint unsigned NOT NULL DEFAULT 0,
    `position_data` blob NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists `group_info` (
    `group_id` bigint unsigned NOT NULL DEFAULT 0,
    `group_name` varchar(255) NOT NULL DEFAULT '',
    `group_avatar` varchar(255) NOT NULL DEFAULT '',
    `group_type` int unsigned NOT NULL DEFAULT 0,
    `group_sub_type` int unsigned NOT NULL DEFAULT 0,
    `server_id` int unsigned NOT NULL DEFAULT 0,
    `manager_list` blob,
    `at_all_per_day` int NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;