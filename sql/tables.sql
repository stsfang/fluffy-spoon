 -- create user table
 CREATE TABLE `tbl_user` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
     `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'user encoded password',
     `avatar` varchar(128)  DEFAULT '' COMMENT 'user avatar',
     `signature` varchar(256) DEFAULT '' COMMENT 'user signature text',
     `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'creating date time',
     `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active timestamp',
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_username` (`user_name`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

 