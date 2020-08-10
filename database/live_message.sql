CREATE TABLE `live_message` (
  `id` bigint(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `msg_type` tinyint(4) NOT NULL COMMENT '消息类型，0控制消息，1文本消息，2图片，3大表情',
  `content` varchar(200) NOT NULL DEFAULT '' COMMENT '消息内容',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '发送者id',
  `receiver_type` tinyint(4) NOT NULL COMMENT '接收者类型，1全部，2at用户，3at消息',
  `receiver_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '接收者id，全部时忽略，at用户时表示用户id，at消息时表示消息id',
  `room_id` int(10) UNSIGNED NOT NULL COMMENT '直播间的id',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB COMMENT='消息表';