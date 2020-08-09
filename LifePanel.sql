-- --------------------------------------------------------
-- Hôte :                        5.196.26.201
-- Version du serveur:           10.3.23-MariaDB-0+deb10u1 - Debian 10
-- SE du serveur:                debian-linux-gnu
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Listage de la structure de la base pour altislife
CREATE DATABASE IF NOT EXISTS `altislife` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `altislife`;

-- Listage de la structure de la table altislife. perms
CREATE TABLE IF NOT EXISTS `perms` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Listage des données de la table altislife.perms : ~3 rows (environ)
/*!40000 ALTER TABLE `perms` DISABLE KEYS */;
REPLACE INTO `perms` (`id`, `name`) VALUES
	(1, 'Fondator'),
	(2, 'Admin'),
	(3, 'Moderator');
/*!40000 ALTER TABLE `perms` ENABLE KEYS */;

-- Listage de la structure de la table altislife. session
CREATE TABLE IF NOT EXISTS `session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `token` text NOT NULL,
  `expiration` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_session_users` (`user`),
  CONSTRAINT `FK_session_users` FOREIGN KEY (`user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=283 DEFAULT CHARSET=utf8;

-- Listage des données de la table altislife.session : ~1 rows (environ)
/*!40000 ALTER TABLE `session` DISABLE KEYS */;
/*!40000 ALTER TABLE `session` ENABLE KEYS */;

-- Listage de la structure de la table altislife. users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login` varchar(50) NOT NULL DEFAULT '',
  `hashPass` text NOT NULL,
  `perms` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login` (`login`),
  KEY `FK_users_perms` (`perms`),
  CONSTRAINT `FK_users_perms` FOREIGN KEY (`perms`) REFERENCES `perms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- Listage des données de la table altislife.users : ~2 rows (environ)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
REPLACE INTO `users` (`id`, `login`, `hashPass`, `perms`) VALUES
	(1, 'Fondator', '$2y$10$cwWtrGHzZulJe/XmCV3ucOyoruOEJwr8MjaimoVL5/TPDGPlX1miq', 1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
