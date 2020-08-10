CREATE TABLE `live_room` (
  `id` int(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '直播间名称',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '直播间描述',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '直播间开启状态',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',

  KEY `name` (`name`)
) ENGINE=InnoDB COMMENT='直播间表';