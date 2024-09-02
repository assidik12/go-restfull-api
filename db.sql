
CREATE DATABASE `go_rest_api`;
USE `go_rest_api`;

CREATE DATABASE `go_rest_api_testing`;

USE `go_rest_api_testing`

CREATE TABLE `category` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) engine=InnoDB;

CREATE TABLE `account` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(300) NOT NULL,
    PRIMARY KEY (`id`)
) engine=InnoDB;

CREATE TABLE `product` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `price` INT NOT NULL,
    `stock` INT NOT NULL,
    `gambar` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `category_id` INT NOT NULL,
    
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
)ENGINE=InnoDB;

CREATE TABLE `transaction_detail` ( 
    `id` INT NOT NULL AUTO_INCREMENT,
    `transaction_id` VARCHAR(300) NOT NULL,
    `product_id` INT NOT NULL,
    `price` INT NOT NULL,
    `quantity` INT NOT NULL,
    PRIMARY KEY (`id`)
) engine=InnoDB;

CREATE TABLE `transaction` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `transaction_detail` VARCHAR(300) NOT NULL,
    `total_price` INT NOT NULL,
    PRIMARY KEY (`id`)
) engine=InnoDB;

ALTER TABLE `transaction` ADD UNIQUE(`transaction_detail`);

ALTER TABLE `transaction`
ADD CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `account` (`id`);


ALTER TABLE `category` ADD COLUMN `stock` INT AFTER `name`;


INSERT INTO `category`(`id`, `name`, `stock`) VALUES(1, 'test', 5);


INSERT INTO `account`(`id`, `username`, `email`, `password`) VALUES(1, 'test', 'test', 'test');

ALTER TABLE `product` ADD COLUMN `img` INT AFTER `description`;

INSERT INTO`product`(`id`, `name`, `price`, `stock`, `description`, `img`, `category_id`) VALUES(1, 'test', 10000, 1, 'test product description', 'test.png', 1);

INSERT INTO `transaction_detail` (`id`, `transaction_id`, `product_id`, `price`, `quantity`) VALUES (2, 1, 1, 50000, 5);



INSERT INTO `transaction`(`id`, `user_id`, `transaction_detail`, `total_price`) VALUES(2, 1, 1, 250000);

SELECT * FROM `transaction` INNER JOIN `transaction_detail` ON `transaction`.`transaction_detail` = `transaction_detail`.`transaction_id` JOIN `product` ON `transaction_detail`.`product_id` = `product`.`id` JOIN `account` ON `transaction`.`user_id` = `account`.`id` WHERE `account`.`id` = 1;