/*
SQLyog Ultimate v11.28 (64 bit)
MySQL - 5.5.60-0ubuntu0.14.04.1-log : Database - mage
表结构初始化
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user`;

/*Table structure for table `user` */

CREATE TABLE `user` (
    `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '账号id',
    `user_type` int(1) NOT NULL DEFAULT '0' COMMENT '账户类型(0:系统账号)',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

/*Data for the table `user` */
insert  into `user`(`user_id`,`user_type`) values (0,0);


SET FOREIGN_KEY_CHECKS = 1;
