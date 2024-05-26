CREATE DATABASE IF NOT EXISTS `devbook`;
USE devbook;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `nick` varchar(50) NOT NULL UNIQUE,
  `email` varchar(50) NOT NULL UNIQUE,
  `password` varchar(50) NOT NULL,
  `createdAt` timestamp DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=INNODB;