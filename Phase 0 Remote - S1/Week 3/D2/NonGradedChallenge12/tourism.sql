/*
 Navicat Premium Data Transfer

 Source Server         : HACTIV8
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : tourism

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 05/09/2023 16:39:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bookings
-- ----------------------------
DROP TABLE IF EXISTS `bookings`;
CREATE TABLE `bookings`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guest_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `check_in_date` date NULL DEFAULT NULL,
  `check_out_date` date NULL DEFAULT NULL,
  `hotel_id` int(11) NULL DEFAULT NULL,
  `total_price` decimal(10, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `hotel_id`(`hotel_id`) USING BTREE,
  CONSTRAINT `bookings_ibfk_1` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bookings
-- ----------------------------
INSERT INTO `bookings` VALUES (1, 'John Doe', '2023-09-10', '2023-09-15', 1, 1500.00);
INSERT INTO `bookings` VALUES (2, 'Alice Smith', '2023-10-05', '2023-10-10', 2, 1800.00);
INSERT INTO `bookings` VALUES (3, 'Bob Johnson', '2023-11-20', '2023-11-25', 3, 1200.00);

-- ----------------------------
-- Table structure for destinations
-- ----------------------------
DROP TABLE IF EXISTS `destinations`;
CREATE TABLE `destinations`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `destination_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `rating` decimal(3, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of destinations
-- ----------------------------
INSERT INTO `destinations` VALUES (1, 'Bali', 'Pulau indah di Indonesia', 'Indonesia', 4.50);
INSERT INTO `destinations` VALUES (2, 'Paris', 'Kota cinta di Prancis', 'Prancis', 4.70);
INSERT INTO `destinations` VALUES (3, 'New York', 'Kota yang tidak pernah tidur di Amerika Serikat', 'Amerika Serikat', 4.30);

-- ----------------------------
-- Table structure for hotels
-- ----------------------------
DROP TABLE IF EXISTS `hotels`;
CREATE TABLE `hotels`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hotel_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `destination_id` int(11) NULL DEFAULT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `rating` decimal(3, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `destination_id`(`destination_id`) USING BTREE,
  CONSTRAINT `hotels_ibfk_1` FOREIGN KEY (`destination_id`) REFERENCES `destinations` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hotels
-- ----------------------------
INSERT INTO `hotels` VALUES (1, 'The Ritz-Carlton Bali', 1, 'Hotel mewah di tepi pantai Bali', 'Jalan Raya Nusa Dua Selatan, Bali', 4.80);
INSERT INTO `hotels` VALUES (2, 'The Eiffel Tower Hotel', 2, 'Hotel dengan pemandangan Menara Eiffel', '123 Rue de Paris, Paris', 4.60);
INSERT INTO `hotels` VALUES (3, 'Times Square Suites', 3, 'Hotel di pusat Times Square', '123 Broadway, New York', 4.20);

SET FOREIGN_KEY_CHECKS = 1;
