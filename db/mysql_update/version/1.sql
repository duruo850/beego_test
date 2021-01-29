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


/*Table structure for table `app_version` */


CREATE TABLE `app_version` (
                               `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                               `version` varchar(32) NOT NULL,
                               `platform` int(1) NOT NULL COMMENT '平台，1 iOS，2 Android',
                               `type` int(1) NOT NULL DEFAULT '0' COMMENT '1 测试版本，2 正式版本',
                               `change_log` text NOT NULL,
                               `is_forced` int(1) NOT NULL DEFAULT '0' COMMENT '0 不强制，1 强制',
                               `dl_path` varchar(256) NOT NULL,
                               `flag` int(1) NOT NULL DEFAULT '1' COMMENT '0 无效，1 有效',
                               `publish_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `app_version` */

insert  into `app_version`(`id`,`version`,`platform`,`type`,`change_log`,`is_forced`,`dl_path`,`flag`,`publish_time`) values (1,'1.1.0',1,1,'',0,'itms-apps://itunes.apple.com/cn/app/%E9%AA%91%E8%AE%B0-%E4%B8%93%E4%B8%9A%E5%8D%95%E8%BD%A6-%E8%87%AA%E8%A1%8C%E8%BD%A6%E9%AA%91%E8%A1%8C%E8%BF%90%E5%8A%A8%E8%BD%A8%E8%BF%B9%E8%AE%B0%E5%BD%95%E8%BD%AF%E4%BB%B6/id596677920?mt=8',1,'2017-11-17 15:49:32'),(2,'12033',2,1,'',0,'http://oz4q5jfbi.bkt.clouddn.com/4024183e-cb63-11e7-9019-027a0cf4beae.1510901279.jpg',1,'2017-11-16 15:49:32');

/*Table structure for table `language` */

CREATE TABLE `language` (
                            `id` int(11) NOT NULL,
                            `desc` varchar(100) NOT NULL,
                            `key` varchar(100) NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `language` */

insert  into `language`(`id`,`desc`,`key`) values (1,'简体中文','zh-Hans'),(2,'英文','en');

/*Table structure for table `user` */

CREATE TABLE `user` (
                        `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '账号id',
                        `user_type` int(1) NOT NULL DEFAULT '0' COMMENT '账户类型(0:系统账号)',
                        `user_name` varchar(32) DEFAULT NULL DEFAULT '' COMMENT '用户名',
                        `password` varchar(64) DEFAULT NULL DEFAULT '' COMMENT '密码',
                        `phone` varchar(13) NOT NULL DEFAULT '' COMMENT '手机号码',
                        `create_time` datetime NOT NULL COMMENT '创建时间',
                        `nick_name` varchar(32) DEFAULT '' COMMENT '姓名',
                        `birthday` date DEFAULT '1990-01-01' COMMENT '生日',
                        `sex` int(1) DEFAULT '0' COMMENT '性别',
                        `email` varchar(64) DEFAULT '' COMMENT '邮箱',
                        `height` int(5) DEFAULT '170' COMMENT '身高',
                        `weight` int(5) DEFAULT '60' COMMENT '体重',
                        `avatar` varchar(128) DEFAULT 'http://avatar.qiteck.net/avatar_default.png' COMMENT '头像',
                        `ftp` int(11) unsigned DEFAULT '100' COMMENT 'ftp',
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `phone` (`phone`) USING BTREE,
                        KEY `user_name` (`user_name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

/*Data for the table `user` */

insert  into `user`(`user_id`,`user_type`,`user_name`,`password`,`phone`,`nick_name`,`birthday`,`sex`,`email`,`height`,`weight`,`avatar`,`ftp`,`create_time`) values (0,0,'test','85777f270ad7cf2a790981bbae3c4e484a1dc55e24a77390d692fbf1cffa12fa','','test',NULL,0,'',170,60,'http://avatar.qiteck.net/avatar_default.png',NULL,'2018-10-18 10:15:11');

/*Table structure for table `authorization` */

CREATE TABLE `authorization` (
                                 `key` varchar(64) NOT NULL DEFAULT '' COMMENT '授权key',
                                 `secret` varchar(64) NOT NULL DEFAULT '' COMMENT '授权秘钥',
                                 `atype` varchar(20) DEFAULT '' COMMENT '授权类型:app, server',
                                 `name` varchar(30) DEFAULT '' COMMENT '授权名称',
                                 `create_time` datetime NOT NULL COMMENT '创建时间',
                                 `user_id` int(11) unsigned DEFAULT NULL,
                                 `valid` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效，1有效，0无效',
                                 PRIMARY KEY (`key`),
                                 KEY `user_id` (`user_id`),
                                 CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='授权表：包括app授权和server授权';

/*Data for the table `authorization` */

insert  into `authorization`(`key`,`secret`,`atype`, `create_time`, `user_id`) values ('424518e4d27b11e8ada274e5f95979ae','d72ac8b653ef83f0da57c394dc743eb0cb674bc7','server','2018-10-18 10:15:11', '0');


/*Table structure for table `user_sns` */
DROP TABLE IF EXISTS `user_sns`;
CREATE TABLE `user_sns` (
                            `user_id` int(11) NOT NULL COMMENT '用户id',
                            `wechat_id` varchar(50) DEFAULT NULL COMMENT 'wechat sns id',
                            `wechat_name` varchar(50) DEFAULT NULL COMMENT 'wechat sns名称',
                            `xiaomi_id` varchar(50) DEFAULT NULL COMMENT 'xioami sns id',
                            `xiaomi_name` varchar(50) DEFAULT NULL COMMENT 'xiaomi sns名称',
                            PRIMARY KEY (`user_id`),
                            KEY `wechat_id` (`wechat_id`),
                            KEY `xiaomi_id` (`xiaomi_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*Data for the table `user_sns` */


/*Table structure for table `whitelist` */
DROP TABLE IF EXISTS `whitelist`;
CREATE TABLE `whitelist` (
                             `phone` varchar(13) NOT NULL COMMENT '手机号码',
                             PRIMARY KEY (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

/*Data for the table `whitelist` */
insert  into `whitelist`(`phone`) values ('11111111111');
