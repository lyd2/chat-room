CREATE TABLE `live_user` (
  `id` int(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT 'username',
  `password` char(32) NOT NULL COMMENT 'password',
  `phone` char(11) NOT NULL DEFAULT '' COMMENT 'phone',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB COMMENT='用户表';