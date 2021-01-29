/*!40101 SET NAMES utf8 */;

CREATE TABLE `db_version` (
    `db_version` int(11) NOT NULL DEFAULT '1' COMMENT '数据库版本'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `db_version` */

insert  into `db_version`(`db_version`) values (1);