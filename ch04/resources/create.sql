CREATE DATABASE IF NOT EXISTS acme;

CREATE TABLE IF NOT EXISTS `acme`.`person` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(100) NOT NULL,
  `phone` CHAR(15) NOT NULL,
  `currency` CHAR(3) NOT NULL,
  `price` DECIMAL(6,2) NOT NULL,
  PRIMARY KEY (`id`));

INSERT INTO `acme`.`person` (`id`, `fullName`, `phone`, `currency`, `price`)
  VALUES ("1", "John", "0123456780", "USD", 100);
INSERT INTO `acme`.`person` (`id`, `fullName`, `phone`, `currency`, `price`)
  VALUES ("2", "Paul", "0123456781", "AUD", 120);
INSERT INTO `acme`.`person` (`id`, `fullName`, `phone`, `currency`, `price`)
  VALUES ("3", "George", "0123456782", "GBP", 150);
INSERT INTO `acme`.`person` (`id`, `fullName`, `phone`, `currency`, `price`)
  VALUES ("4", "Ringo", "0123456783", "EUR", 110);

