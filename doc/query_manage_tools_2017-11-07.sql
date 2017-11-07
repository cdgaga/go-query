# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 192.168.10.10 (MySQL 5.5.50-0ubuntu0.14.04.1)
# Database: query_manage_tools
# Generation Time: 2017-11-07 08:09:10 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table act_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `act_log`;

CREATE TABLE `act_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `act_type` varchar(50) NOT NULL DEFAULT '',
  `act_title` varchar(2500) NOT NULL DEFAULT '',
  `act_content` text,
  `act_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table db_config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `db_config`;

CREATE TABLE `db_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `db_name` varchar(100) NOT NULL DEFAULT '',
  `db_type` varchar(50) NOT NULL COMMENT 'db type : mysql,redis and so on',
  `db_host` varchar(100) NOT NULL DEFAULT '',
  `db_user` varchar(100) NOT NULL,
  `db_password` varchar(100) NOT NULL DEFAULT '',
  `db_port` varchar(10) NOT NULL DEFAULT '',
  `charset` varchar(50) NOT NULL,
  `sub_db` varchar(2500) DEFAULT NULL COMMENT 'sub db name',
  `tables` text COMMENT 'table name',
  `create_time` datetime NOT NULL COMMENT 'create time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `db_config` WRITE;
/*!40000 ALTER TABLE `db_config` DISABLE KEYS */;

INSERT INTO `db_config` (`id`, `db_name`, `db_type`, `db_host`, `db_user`, `db_password`, `db_port`, `charset`, `sub_db`, `tables`, `create_time`)
VALUES
	(1,'test','mysql','127.0.0.1','','pass','3306','utf8',NULL,NULL,'0000-00-00 00:00:00');

/*!40000 ALTER TABLE `db_config` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table role_db_relation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `role_db_relation`;

CREATE TABLE `role_db_relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `db_id` int(11) NOT NULL,
  `is_query_all` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1 part 2 all',
  `query_table` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'Login Name',
  `user_pwd` varchar(100) NOT NULL DEFAULT '' COMMENT 'Login password',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT 'Login Status 1 enable 2 disable',
  `login_last_time` datetime NOT NULL,
  `login_last_ip` varchar(100) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `is_super_admin` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '1 super admin 0 none',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`user_id`, `user_name`, `user_pwd`, `status`, `login_last_time`, `login_last_ip`, `create_time`, `is_super_admin`)
VALUES
	(1,'admin','e99a18c428cb38d5f260853678922e03',1,'2017-07-01 12:31:20','127.0.0.1','0000-00-00 00:00:00',1),
	(2,'xxx','e99a18c428cb38d5f260853678922e03',1,'2017-07-01 12:31:21','11','0000-00-00 00:00:00',0),
	(3,'xxxx','e99a18c428cb38d5f260853678922e03',1,'2017-07-01 12:31:21','11','0000-00-00 00:00:00',1);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_favorites
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_favorites`;

CREATE TABLE `user_favorites` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `query_db_type` varchar(50) DEFAULT NULL,
  `query_sql` varchar(2500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
