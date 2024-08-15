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



DESCRIBE `category`


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

ALTER TABLE `customer` ADD COLUMN `phone` VARCHAR(255)  AFTER `name`;
ALTER TABLE `product` ADD COLUMN `description` VARCHAR(255)  AFTER `stock`;