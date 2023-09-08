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

-- Trigger untuk memastikan setiap buku hanya dapat dipinjamkan kepada satu anggota dalam satu waktu:
DELIMITER //
CREATE TRIGGER prevent_multiple_borrowers
BEFORE INSERT ON lending_transaction
FOR EACH ROW
BEGIN
  DECLARE book_count INT;
  SELECT COUNT(*) INTO book_count FROM lending_transaction WHERE isbn = NEW.isbn AND date_of_return IS NULL;
  IF book_count > 0 THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Buku ini sedang dipinjam oleh anggota lain.';
  END IF;
END;
//
DELIMITER ;

-- Trigger untuk memastikan anggota dapat meminjamkan tidak lebih dari 5 buku sekaligus:
DELIMITER //
CREATE TRIGGER prevent_excessive_borrowing
BEFORE INSERT ON lending_transaction
FOR EACH ROW
BEGIN
  DECLARE borrowed_count INT;
  SELECT COUNT(*) INTO borrowed_count FROM lending_transaction WHERE member_id = NEW.member_id AND date_of_return IS NULL;
  IF borrowed_count >= 5 THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Anda tidak dapat meminjam lebih dari 5 buku sekaligus.';
  END IF;
END;
//
DELIMITER ;

-- Insert Data to Table Book
INSERT INTO `small_library_system`.`book` (`isbn`, `title`, `author`, `publisher`, `year_of_publication`) 
VALUES 
('1292-9598-3449', 'Harry Potter and the Sorcerers Stone', 'J.K. Rowling', 'Andi', 2020),
('5784-5421-2352', 'To Kill a Mockingbird', 'Harper Lee', 'Andi', 2019),
('4567-9999-2312', 'The Great Gatsby', 'F. Scott Fitzgerald', 'Andi', 2018);

-- Insert Data to Table Member
INSERT INTO `small_library_system`.`member` (`frist_name`, `last_name`, `date_of_birth`, `address`)
VALUES 
('John', 'Doe', '1990-01-01', 'Jl. Sudirman No. 1'),
('Jane', 'Doe', '1991-01-01', 'Jl. Sudirman No. 2'),
('James', 'Doe', '1992-01-01', 'Jl. Sudirman No. 3');

-- Insert Data to Table Lending Transaction
INSERT INTO `small_library_system`.`lending_transaction` (`member_id`, `isbn`, `date_of_lending`, `date_of_return`, `condition_at_return`)
VALUES 
(1, '1292-9598-3449', '2020-01-01', '2020-01-08', 'Good'),
(2, '5784-5421-2352', '2020-01-01', '2020-01-08', 'Good'),
(3, '4567-9999-2312', '2020-01-01', '2020-01-08', 'Lost');