/*
 Navicat Premium Data Transfer

 Source Server         : gong
 Source Server Type    : MySQL
 Source Server Version : 80020 (8.0.20)
 Source Host           : localhost:3306
 Source Schema         : callersystem

 Target Server Type    : MySQL
 Target Server Version : 80020 (8.0.20)
 File Encoding         : 65001

 Date: 30/12/2022 10:18:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, '2022-12-26 20:10:15.000', '2022-12-26 20:10:19.000', NULL, 'gsy', '123456');

-- ----------------------------
-- Table structure for visitors
-- ----------------------------
DROP TABLE IF EXISTS `visitors`;
CREATE TABLE `visitors`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `visitor_name` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sex` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `phone` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `visit_id` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitors
-- ----------------------------
INSERT INTO `visitors` VALUES (1, '2022-12-26 22:18:32.834', '2022-12-26 22:18:32.834', NULL, 'kkk', '女', '22222222222', '222222222222222222', 'lllll');
INSERT INTO `visitors` VALUES (2, '2022-12-26 22:20:08.140', '2022-12-26 22:20:08.140', NULL, 'lll', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (3, '2022-12-26 22:20:08.973', '2022-12-26 22:20:08.973', NULL, 'lll', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (4, '2022-12-26 22:20:09.758', '2022-12-26 22:20:09.758', NULL, 'lll', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (5, '2022-12-27 13:36:11.662', '2022-12-27 13:36:11.662', NULL, 'dddd', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (6, '2022-12-27 13:53:02.345', '2022-12-27 13:53:02.345', NULL, 'dddd', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (7, '2022-12-27 14:25:16.515', '2022-12-27 14:25:16.515', NULL, 'ffff', '男', '22222222222', '222222222222222222', '进入');
INSERT INTO `visitors` VALUES (8, '2022-12-27 15:16:38.152', '2022-12-27 15:16:38.152', NULL, 'dddd', '女', '11111111111', '111111111111111111', '事由');
INSERT INTO `visitors` VALUES (9, '2022-12-27 15:23:07.668', '2022-12-27 15:23:07.668', NULL, 'gsy', '男', '55555555555', '555555555555555555', 'tt');
INSERT INTO `visitors` VALUES (10, '2022-12-27 15:27:33.774', '2022-12-27 15:27:33.774', NULL, 'zll', '男', '66666666666', '666666666666666666', 'ssssdsd');
INSERT INTO `visitors` VALUES (11, '2022-12-27 15:29:09.079', '2022-12-27 15:29:09.079', NULL, 'ldn', '女', '99999999999', '999999999999999999', 'kkkkk');
INSERT INTO `visitors` VALUES (12, '2022-12-27 15:30:49.913', '2022-12-27 15:30:49.913', NULL, 'qyk', '男', '77777777777', '777777777777777777', 'lllll');
INSERT INTO `visitors` VALUES (13, '2022-12-27 15:38:51.084', '2022-12-27 15:38:51.084', NULL, 'zy', '男', '00000000000', '000000000000000000', 'llll');
INSERT INTO `visitors` VALUES (14, '2022-12-27 15:39:11.141', '2022-12-27 15:39:11.141', NULL, 'zy', '男', '00000000000', '000000000000000000', 'llll');
INSERT INTO `visitors` VALUES (15, '2022-12-27 15:41:49.352', '2022-12-27 15:41:49.352', NULL, 'gsy', '男', '55555555555', '555555555555555555', 'lllll');
INSERT INTO `visitors` VALUES (16, '2022-12-27 15:43:20.828', '2022-12-27 15:43:20.828', NULL, 'gsy', '男', '55555555555', '555555555555555555', 'llllll');
INSERT INTO `visitors` VALUES (17, '2022-12-28 15:24:15.416', '2022-12-28 15:24:15.416', NULL, 'lllll', '女', '33333333333', '333333333333333333', 'sdsdsds');
INSERT INTO `visitors` VALUES (18, '2022-12-29 14:00:45.808', '2022-12-29 14:00:45.808', NULL, 'leo', '女', '77777777777', '888888888888888888', '进入');
INSERT INTO `visitors` VALUES (19, '2022-12-29 16:32:43.113', '2022-12-29 16:32:43.113', NULL, 'zv', '男', '66666666666', '123456789123456789', '进入');
INSERT INTO `visitors` VALUES (20, '2022-12-29 16:34:41.807', '2022-12-29 16:34:41.807', NULL, 'zv', '男', '66666666666', '123456789123456789', '进入');
INSERT INTO `visitors` VALUES (21, '2022-12-30 08:40:19.557', '2022-12-30 08:40:19.557', NULL, 'ppp', '男', '77777777777', '888888888888888888', '进入');
INSERT INTO `visitors` VALUES (22, '2022-12-30 08:45:53.995', '2022-12-30 08:45:53.995', NULL, 'ppp', '女', '77777777777', '888888888888888888', '111');

SET FOREIGN_KEY_CHECKS = 1;
