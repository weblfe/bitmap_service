CREATE TABLE `user_day_login`
(
    `id`                bigint(20)   NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `date`              date         NOT NULL COMMENT '日期',
    `uid`               bigint(20)   NOT NULL COMMENT 'uid',
    `count`             int(10)      NOT NULL COMMENT '登陆次数',
    `is_anchor`         smallint(1)  NOT NULL COMMENT '是否认证主播(0:非主播,1:主播)',
    `is_newer`          smallint(1)  NOT NULL COMMENT '是否新用户(0:老用户,1:新用户)',
    `user_class`        smallint(1)  NOT NULL COMMENT '用户类型(0:非付费用户,1:付费用户)',
    `channel`           varchar(100) NOT NULL COMMENT '用户所属渠道名',
    `reg_channel`       varchar(100) NOT NULL COMMENT '用户注册渠道名',
    `devices_id`        varchar(100) NOT NULL COMMENT '用户登录设备号',
    `start_activity_at` datetime     NOT NULL COMMENT '当天首次活跃时间点',
    `last_activity_at`  datetime     NOT NULL COMMENT '当天最后一次活跃时间点',
    `created_at`        datetime     NOT NULL COMMENT '创建时间',
    `updated_at`        datetime     NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_login` (`date`, `uid`),
    KEY `idx_user_query` (`is_newer`, `is_anchor`, `channel`),
    KEY `idx_reg_channel` (`reg_channel`, `is_newer`, `is_anchor`),
    KEY `idx_uid_devices_id` (`uid`, `devices_id`, `is_anchor`)
) ENGINE = MyISAM
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='(日访问)用户每日登陆记录';