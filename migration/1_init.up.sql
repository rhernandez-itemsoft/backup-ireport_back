/*
 Navicat Premium Data Transfer

 Source Server         : Maria DB localhost
 Source Server Type    : MySQL
 Source Server Version : 100412
 Source Host           : localhost:3306
 Source Schema         : ireport

 Target Server Type    : MySQL
 Target Server Version : 100412
 File Encoding         : 65001

 Date: 10/06/2020 23:33:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for datasource_auth
-- ----------------------------
CREATE TABLE `datasource_auths`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `bearer` text NULL,
  `username` varchar(30) NULL DEFAULT '',
  `password` varchar(30) NULL DEFAULT '',
  `apiSource` varchar(15) NULL DEFAULT '' COMMENT 'Header, QueryParams',
  `apiKey` varchar(30) NULL DEFAULT '',
  `apiValue` text NULL,
  `datasource_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `datasourceId_idx`(`datasource_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for datasource_param
-- ----------------------------
CREATE TABLE `datasource_params`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `key` varchar(30) NOT NULL,
  `type` varchar(30) NOT NULL,
  `datasource_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `datasourceId_idx`(`datasource_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for datasource
-- ----------------------------
CREATE TABLE `datasources`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `description` varchar(255) NULL DEFAULT '',
  `method` varchar(5) NOT NULL DEFAULT '',
  `endpoint` text NOT NULL,
  `authType` varchar(20) NOT NULL DEFAULT '',
  `accept` varchar(50) NOT NULL DEFAULT '',
  `contentType` varchar(50) NOT NULL DEFAULT '',
  `response` text NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;




SET FOREIGN_KEY_CHECKS = 1;
