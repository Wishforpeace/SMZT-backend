CREATE DATABASE `SMZT`;
use SMZT;
-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `created_at`          varchar(255) DEFAULT NULL,
    `updated_at`          varchar(255)  DEFAULT NULL,
    `deleted_at`          varchar(255)  ,
    `id`            int(11)     NOT NULL AUTO_INCREMENT,
    `student_id`         varchar(35) NOT NULL,
    `hash_password` varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `student_id` (`student_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;
