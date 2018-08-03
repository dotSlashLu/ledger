CREATE TABLE `expense_class` (
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(32) NOT NULL,
	`level` TINYINT NOT NULL DEFAULT 0,
	`parent` INT DEFAULT NULL,
	`create_time` DATETIME DEFAULT NULL
) CHARSET=utf8;

INSERT INTO `expense_class` VALUES 
(1,'交通',0,NULL,'2018-07-27 15:40:32'),
(5,'公交',1,1,'2018-07-27 15:43:27'),
(6,'打车',1,1,'2018-07-27 15:43:27'),
(2,'其它',0,NULL,'2018-07-27 15:40:47'),
(3,'人情',0,NULL,'2018-07-27 15:41:24'),
(4,'宠物',0,NULL,'2018-07-27 15:41:30'),
(7,'测试',0,NULL,'2018-07-27 15:43:27'),
(8,'测试儿子',1,7,'2018-07-27 15:43:27'),
(9,'商超',0,NULL,'2018-07-27 15:43:27'),
(10, '餐饮',0,NULL,'2018-07-27 15:43:27'),
(11,'饭店',1,10,'2018-07-27 15:43:27'),
(12,'小吃',1,10,'2018-07-27 15:43:27'),
(13,'饮品',1,10,'2018-07-27 15:43:27'),
(14,'外卖',1,10,'2018-07-27 15:43:27');

CREATE TABLE `expense` (
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`title` VARCHAR(32) DEFAULT NULL,
	`cost` FLOAT(10,3) NOT NULL,
	`class` INT DEFAULT NULL,
	`sub_class` INT DEFAULT NULL,
	`remark` VARCHAR(512) DEFAULT "",
	`create_time` DATETIME DEFAULT NULL,
	`uid` INT NOT NULL,

	KEY `create time`(`create_time`),
	KEY `class`(`class`),
	KEY `sub class`(`sub_class`),
	KEY `by user`(`uid`)
) CHARSET=utf8;

CREATE TABLE `user` (
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`username` VARCHAR(32) NOT NULL,
	`password` VARCHAR(128) NOT NULL,
	`create_time` DATETIME DEFAULT NULL,

	KEY `validate` (`username`, `password`)
) CHARSET=utf8;