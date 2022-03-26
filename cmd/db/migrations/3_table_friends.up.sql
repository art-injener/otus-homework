CREATE TABLE `friends` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `user_id` int NOT NULL COMMENT 'id пользователя',
                           `friend_id` int NOT NULL COMMENT 'id друга',
                           `accept` tinyint(1) DEFAULT '1' COMMENT 'флаг подтверждения обоюдной дружбы',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `friends_id_uindex` (`id`),
                           KEY `friends_users_id_fk` (`user_id`),
                           KEY `friends_users_id_fk_2` (`friend_id`),
                           CONSTRAINT `friends_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
                           CONSTRAINT `friends_users_id_fk_2` FOREIGN KEY (`friend_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='таблица друзей ';