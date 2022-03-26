CREATE TABLE `logins_info` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `email` varchar(100) DEFAULT NULL,
                               `login` varchar(100) NOT NULL,
                               `password` varchar(100) NOT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `logins_info_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;