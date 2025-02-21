 CREATE DATABASE IF NOT EXISTS `cart`  
    DEFAULT CHARACTER SET = 'utf8mb4';;

CREATE DATABASE IF NOT EXISTS `checkout`
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS `order`
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS `payment`
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS `product`
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS `user`
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS `casbin`
    DEFAULT CHARACTER SET = 'utf8mb4';

USE `product`;
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `description` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_category_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `category` WRITE;
INSERT INTO `category` VALUES (1,'2023-12-06 15:05:06.000','2023-12-06 15:05:06.000','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06.000','2023-12-06 15:05:06.000','Sticker','Sticker'),(5,'2025-02-20 09:06:47.320','2025-02-20 09:06:47.320','toy','toy');
UNLOCK TABLES;

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `description` longtext,
  `picture` longtext,
  `price` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `product` WRITE;
INSERT INTO `product` VALUES (1,'2023-12-06 15:26:19.000','2023-12-09 22:29:10.000','Notebook','The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ','/static/image/notebook.jpeg',9.9),(2,'2023-12-06 15:26:19.000','2023-12-09 22:29:10.000','Mouse-Pad','The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ','/static/image/mouse-pad.jpeg',8.8),(3,'2023-12-06 15:26:19.000','2023-12-09 22:31:20.000','T-Shirt','The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.','/static/image/t-shirt.jpeg',6.6),(4,'2023-12-06 15:26:19.000','2023-12-09 22:31:20.000','T-Shirt','The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.','/static/image/t-shirt-1.jpeg',2.2),(5,'2023-12-06 15:26:19.000','2023-12-09 22:32:35.000','Sweatshirt','The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.','/static/image/sweatshirt.jpeg',1.1),(6,'2023-12-06 15:26:19.000','2023-12-09 22:31:20.000','T-Shirt','The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.','/static/image/t-shirt-2.jpeg',1.8),(7,'2023-12-06 15:26:19.000','2023-12-09 22:31:20.000','mascot','The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.','/static/image/logo.jpg',4.8),(10,'2025-02-20 09:11:04.607','2025-02-20 09:11:04.607','cat toy','cat toy','',100);
UNLOCK TABLES;

DROP TABLE IF EXISTS `product_category`;

CREATE TABLE `product_category` (
  `category_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  PRIMARY KEY (`category_id`,`product_id`),
  KEY `fk_product_category_product` (`product_id`),
  CONSTRAINT `fk_product_category_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `fk_product_category_product` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `product_category` WRITE;
INSERT INTO `product_category` VALUES (2,1),(2,2),(1,3),(1,4),(1,5),(1,6),(2,7),(5,10);
UNLOCK TABLES;


USE `casbin`;
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `casbin_rule` WRITE;
INSERT INTO `casbin_rule` VALUES (13,'g','admin','user','','','',''),(2,'p','user','/cart','GET','','',''),(1,'p','user','/cart','POST','','',''),(6,'p','user','/checkout','GET','','',''),(8,'p','user','/checkout/result','GET','','',''),(7,'p','user','/checkout/waiting','POST','','',''),(3,'p','user','/order','GET','','','');
UNLOCK TABLES;

USE `user`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password_hashed` longtext,
  `role` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `user` WRITE;
INSERT INTO `user` VALUES (1,'2025-02-21 01:10:13.178','2025-02-21 01:10:13.178','admin@admin.com','$2a$10$6vHL4gpdX1N849W1FVMda.nuIFIgmVXdo9JQVM5T91Od5GO/FApl2','admin');
UNLOCK TABLES;