/*
 Navicat Premium Data Transfer

 Source Server         : docker
 Source Server Type    : MySQL
 Source Server Version : 80019 (8.0.19)
 Source Host           : 127.0.0.1:3306
 Source Schema         : cloud_storage

 Target Server Type    : MySQL
 Target Server Version : 80019 (8.0.19)
 File Encoding         : 65001

 Date: 21/12/2022 20:16:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for storage_repository_pool
-- ----------------------------
DROP TABLE IF EXISTS `storage_repository_pool`;
CREATE TABLE `storage_repository_pool` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `hash` varchar(32) DEFAULT NULL COMMENT '文件的唯一标识',
  `name` varchar(255) DEFAULT NULL,
  `ext` varchar(30) DEFAULT NULL COMMENT '文件扩展名',
  `size` int DEFAULT NULL COMMENT '文件大小',
  `path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of storage_repository_pool
-- ----------------------------
BEGIN;
INSERT INTO `storage_repository_pool` (`id`, `identity`, `hash`, `name`, `ext`, `size`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '58f6100b-8f19-4853-a4ae-79218847e6a7', 'e71e6711e0937fb71605e1cfb61bf495', 'aaaa.jpeg', '.jpeg', 45673, 'http://images.caixiaoxin.cn/go-cloud-storage/aaaa.jpeg', '2022-12-20 23:46:20', '2022-12-20 23:46:20', NULL);
INSERT INTO `storage_repository_pool` (`id`, `identity`, `hash`, `name`, `ext`, `size`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '62f8fb75-e62c-4de0-b0a3-0ba1236d05ba', 'FoJeETy1AW5A-dh13H1qtKNcQRcp', 'tron-jrpc1.jpg', '.jpg', 60972, 'http://images.caixiaoxin.cn/go-cloud-storage/tron-jrpc1.jpg', '2022-12-21 18:23:01', '2022-12-21 18:23:01', NULL);
COMMIT;

-- ----------------------------
-- Table structure for storage_share
-- ----------------------------
DROP TABLE IF EXISTS `storage_share`;
CREATE TABLE `storage_share` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL COMMENT '公共池中的唯一标识',
  `user_repository_identity` varchar(36) DEFAULT NULL COMMENT '用户池子中的唯一标识',
  `expired_time` int DEFAULT NULL COMMENT '失效时间，单位秒, 【0-永不失效】',
  `click_num` int DEFAULT '0' COMMENT '点击次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of storage_share
-- ----------------------------
BEGIN;
INSERT INTO `storage_share` (`id`, `identity`, `user_identity`, `repository_identity`, `user_repository_identity`, `expired_time`, `click_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, '0f198b2a-6ed5-4261-a2c8-445e4e6e84e7', 'e6ab9ffc-887f-4935-81db-65e9d308ac64', '58f6100b-8f19-4853-a4ae-79218847e6a7', '', 10000, 4, '2022-12-21 14:54:22', '2022-12-21 14:54:22', NULL);
COMMIT;

-- ----------------------------
-- Table structure for storage_user
-- ----------------------------
DROP TABLE IF EXISTS `storage_user`;
CREATE TABLE `storage_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(60) DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of storage_user
-- ----------------------------
BEGIN;
INSERT INTO `storage_user` (`id`, `identity`, `name`, `password`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'abc111', 'abc', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'abc@qq.com', '2022-12-20 16:14:56', '2022-12-20 16:14:58', NULL);
INSERT INTO `storage_user` (`id`, `identity`, `name`, `password`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 'ac7712a1-4092-494e-8908-c0015158af9b', 'test1222', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'aa111@qq.com', '2022-12-20 20:48:55', '2022-12-20 20:48:55', NULL);
INSERT INTO `storage_user` (`id`, `identity`, `name`, `password`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'e6ab9ffc-887f-4935-81db-65e9d308ac64', 'test1', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '1752676696@qq.com', '2022-12-20 21:03:29', '2022-12-20 21:03:29', NULL);
COMMIT;

-- ----------------------------
-- Table structure for storage_user_repository
-- ----------------------------
DROP TABLE IF EXISTS `storage_user_repository`;
CREATE TABLE `storage_user_repository` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL,
  `ext` varchar(255) DEFAULT NULL COMMENT '文件或文件夹类型',
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of storage_user_repository
-- ----------------------------
BEGIN;
INSERT INTO `storage_user_repository` (`id`, `identity`, `user_identity`, `parent_id`, `repository_identity`, `ext`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, '2ba3348e-2581-4610-936e-e11a873df9d2', 'e6ab9ffc-887f-4935-81db-65e9d308ac64', 0, '58f6100b-8f19-4853-a4ae-79218847e6a7', '.jpeg', 'aaaa.jpeg', '2022-12-20 23:51:54', '2022-12-21 00:28:48', '2022-12-21 14:31:01');
INSERT INTO `storage_user_repository` (`id`, `identity`, `user_identity`, `parent_id`, `repository_identity`, `ext`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'e821471d-e0ae-487a-83db-053ced895c7b', 'e6ab9ffc-887f-4935-81db-65e9d308ac64', 0, '', '', '音乐', '2022-12-21 01:13:13', '2022-12-21 01:13:13', NULL);
INSERT INTO `storage_user_repository` (`id`, `identity`, `user_identity`, `parent_id`, `repository_identity`, `ext`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, '2ba3348e-2581-4610-936e-e11a873df9d3', 'e6ab9ffc-887f-4935-81db-65e9d308ac64', 9, '58f6100b-8f19-4853-a4ae-79218847e6a7', '.jpeg', 'aaaa.jpeg', '2022-12-20 23:51:54', '2022-12-21 14:37:13', NULL);
INSERT INTO `storage_user_repository` (`id`, `identity`, `user_identity`, `parent_id`, `repository_identity`, `ext`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, '1990c4c2-2416-4643-8bcf-fd4e67107b00', 'e6ab9ffc-887f-4935-81db-65e9d308ac64', 0, '58f6100b-8f19-4853-a4ae-79218847e6a7', '.jpeg', 'aaaa.jpeg', '2022-12-21 15:26:12', '2022-12-21 15:26:12', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
