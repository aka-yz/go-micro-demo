CREATE DATABASE `entry_task_db`;

CREATE TABLE `entry_task_db`.`activity_info_tab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `activity_name` varchar(20) NOT NULL DEFAULT '' COMMENT '活动名称',
  `activity_start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动开始时间',
  `activity_end_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动结束时间',
  `activity_address` varchar(128) NOT NULL DEFAULT '' COMMENT '活动地点',
  `activity_description` varchar(256) NOT NULL DEFAULT '' COMMENT '活动详细描述',
  `created_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人 id|关联user_info_tab.id',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_start_time` (`activity_start_time`),
  KEY `idx_end_time` (`activity_end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动信息数据表|yue.zhang';

CREATE TABLE `entry_task_db`.`activity_user_associate_tab` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
 `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 id',
 `activity_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动 id|关联activity_info_tab.id',
 `apply_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '报名时间',
 `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
 `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
 PRIMARY KEY (`id`),
 KEY `idx_user` (`user_id`) USING BTREE,
 KEY `idx_activity` (`activity_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动用户关联数据表|yue.zhang';

CREATE TABLE `entry_task_db`.`activity_user_comment_tab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 id|关联user_info_tab.id',
  `activity_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动 id|关联activity_info_tab.id',
  `comment` varchar(256) NOT NULL DEFAULT '' COMMENT '评论',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_activity` (`activity_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动用户评论数据表|yue.zhang';

CREATE TABLE `entry_task_db`.`classification_activity_tab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `activity_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动 id|关联activity_info_tab.id',
  `classification_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动分类 id|关联classification_info_tab.id',
  `created_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人 id|关联user_info_tab.id',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_classify` (`classification_id`) USING BTREE,
  KEY `idx_activity` (`activity_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动分类关联信息数据表|yue.zhang';

CREATE TABLE `entry_task_db`.`classification_info_tab` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
 `classification_name` varchar(20) NOT NULL DEFAULT '' COMMENT '分类名称',
 `created_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人 id|关联user_info_tab.id',
 `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
 `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类信息数据表|yue.zhang';

CREATE TABLE `entry_task_db`.`user_info_tab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像链接',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT '密码（md5 加密）',
  `role` tinyint(3) NOT NULL DEFAULT '0' COMMENT '角色|0-普通用户;1-运营',
  `login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次登录时间',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息数据表|yue.zhang';