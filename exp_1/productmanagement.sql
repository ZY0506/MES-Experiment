/*
 Navicat Premium Dump SQL

 Source Server         : 1
 Source Server Type    : MySQL
 Source Server Version : 80043 (8.0.43)
 Source Host           : localhost:3306
 Source Schema         : productmanagement

 Target Server Type    : MySQL
 Target Server Version : 80043 (8.0.43)
 File Encoding         : 65001

 Date: 13/10/2025 21:15:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '产品名称',
  `category` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '种类',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '编号',
  `price` decimal(10, 2) NOT NULL COMMENT '价格',
  `quantity` int NOT NULL DEFAULT 0 COMMENT '数量',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code` ASC) USING BTREE,
  CONSTRAINT `chk_quantity` CHECK (`quantity` >= 0)
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '产品信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, '影驰 RTX 4070 Ti 星曜', '显卡', 'GPU-001', 5999.00, 12);
INSERT INTO `products` VALUES (2, 'LG 27GP850-B 27寸 NanoIPS', '显示器', 'MON-001', 2199.00, 8);
INSERT INTO `products` VALUES (3, '罗技 G502 HERO', '鼠标', 'MOU-001', 299.00, 35);
INSERT INTO `products` VALUES (4, 'Intel i7-13700KF', 'CPU', 'CPU-001', 2599.00, 20);
INSERT INTO `products` VALUES (5, '三星 980 PRO 1TB NVMe', '硬盘', 'SSD-001', 699.00, 42);
INSERT INTO `products` VALUES (6, '华硕 TUF RTX 4060 Ti 8G', '显卡', 'GPU-002', 3299.00, 15);
INSERT INTO `products` VALUES (7, 'AOC Q27G2S 27寸 2K 155Hz', '显示器', 'MON-002', 1399.00, 10);
INSERT INTO `products` VALUES (8, '雷蛇 毒蝰终极版', '鼠标', 'MOU-002', 499.00, 28);
INSERT INTO `products` VALUES (9, 'AMD Ryzen 7 7800X3D', 'CPU', 'CPU-002', 2999.00, 18);
INSERT INTO `products` VALUES (10, '西数 SN850X 2TB NVMe', '硬盘', 'SSD-002', 1199.00, 22);

SET FOREIGN_KEY_CHECKS = 1;
