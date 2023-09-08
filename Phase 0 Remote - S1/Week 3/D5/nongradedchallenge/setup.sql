-- Create Database Small Libary System
CREATE DATABASE `small_library_system`;

-- Use Database Small Libary System
USE `small_library_system`;

-- Create Tables Book
CREATE TABLE `book` (
  `isbn` varchar(15) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `publisher` varchar(255) DEFAULT NULL,
  `year_of_publication` year DEFAULT NULL,
  PRIMARY KEY (`isbn`)
);

-- Create Tables Member
CREATE TABLE `member` (
  `member_id` int NOT NULL AUTO_INCREMENT,
  `frist_name` varchar(15) DEFAULT NULL,
  `last_name` varchar(15) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`member_id`)
);

-- Create Tables Lending Transaction
CREATE TABLE `lending_transaction` (
  `transaction_id` int NOT NULL AUTO_INCREMENT,
  `member_id` int DEFAULT NULL,
  `isbn` varchar(255) DEFAULT NULL,
  `date_of_lending` date DEFAULT NULL,
  `date_of_return` date DEFAULT NULL,
  `condition_at_return` enum('Good','Damage','Lost') DEFAULT NULL,
  PRIMARY KEY (`transaction_id`),
  CONSTRAINT `member_id_fk` FOREIGN KEY (`member_id`) REFERENCES `member` (`member_id`),
  CONSTRAINT `isbn_fk` FOREIGN KEY (`isbn`) REFERENCES `book` (`isbn`)
);