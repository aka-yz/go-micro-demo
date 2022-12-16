CREATE DATABASE `demo_db`;

CREATE TABLE `demo_db`.`user_info_tab` (
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