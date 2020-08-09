/*
 Navicat Premium Data Transfer

 Source Server         : local_mysql
 Source Server Type    : MySQL
 Source Server Version : 50540
 Source Host           : 127.0.0.1:3306
 Source Schema         : todolist

 Target Server Type    : MySQL
 Target Server Version : 50540
 File Encoding         : 65001

 Date: 09/08/2020 22:35:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for todo_models
-- ----------------------------
DROP TABLE IF EXISTS `todo_models`;
CREATE TABLE `todo_models`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `completed` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_todo_models_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of todo_models
-- ----------------------------
INSERT INTO `todo_models` VALUES (1, '2020-08-02 11:02:13', '2020-08-02 11:27:18', NULL, '踢足球', 0);
INSERT INTO `todo_models` VALUES (2, '2020-08-02 11:06:11', '2020-08-02 11:06:11', NULL, '打篮球', 0);
INSERT INTO `todo_models` VALUES (3, '2020-08-02 11:28:11', '2020-08-02 11:28:11', NULL, '打羽毛球', 0);
INSERT INTO `todo_models` VALUES (4, '2020-08-02 11:28:26', '2020-08-02 11:28:26', NULL, 'lol', 0);
INSERT INTO `todo_models` VALUES (5, '2020-08-02 11:28:33', '2020-08-02 11:28:33', NULL, '王者荣耀', 0);

SET FOREIGN_KEY_CHECKS = 1;
