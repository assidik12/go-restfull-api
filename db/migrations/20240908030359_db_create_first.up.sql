-- Active: 1722607424659@@127.0.0.1@3306@go_rest_api

CREATE TABLE `category` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `stock` INT NOT NULL,
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



