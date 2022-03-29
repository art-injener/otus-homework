CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `login_id` int NOT NULL,
                         `name` varchar(50) DEFAULT NULL,
                         `age` int DEFAULT NULL,
                         `sex` char(1) DEFAULT NULL,
                         `hobby` text,
                         `city` varchar(100) DEFAULT NULL,
                         `surname` varchar(50) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `users_id_uindex` (`id`),
                         KEY `users_logins_info_id_fk` (`login_id`),
                         CONSTRAINT `users_logins_info_id_fk` FOREIGN KEY (`login_id`) REFERENCES `logins_info` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='анкеты пользователей социальной сети';


