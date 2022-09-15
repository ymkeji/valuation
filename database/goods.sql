-- MySQL dump 10.13  Distrib 8.0.30, for macos12.4 (arm64)
--
-- Host: 180.76.246.37    Database: valuation
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goods`
--

DROP TABLE IF EXISTS `goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `type` varchar(10) NOT NULL DEFAULT '',
  `unit` varchar(10) NOT NULL DEFAULT '',
  `price` decimal(10,2) NOT NULL,
  `tariff` float unsigned NOT NULL DEFAULT '0',
  `alias` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `索引_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods`
--

LOCK TABLES `goods` WRITE;
/*!40000 ALTER TABLE `goods` DISABLE KEYS */;
INSERT INTO `goods` VALUES (1,'豆腐','优','斤',30.24,0,'doufu'),(2,'五花肉','优','斤',15.80,0,'wuhuarou'),(3,'茄子','优','斤',1.20,0,'jiazi'),(4,'白萝卜','优','斤',1.40,0,'bailuobo'),(5,'狗肉','优','斤',5.00,0,'gourou'),(6,'油饼','优','个',0.50,0,'youbing'),(7,'大蒜','优','斤',0.50,0,'dasuan'),(8,'肉丝','优','斤',12.00,0,'rousi'),(9,'排骨','优','斤',15.30,0,'paigu'),(10,'圆红','优','斤',22.30,0,'yuanhong'),(11,'羊肉','优','斤',33.40,0,'yangrou'),(12,'花甲','优','斤',122.00,0,'huajia'),(13,'虾尾','优','件',1.20,0,'xiawei'),(14,'青笋','优','斤',123.00,0,'qingsun'),(15,'肉片子','优','斤',3.43,0,'roupianzi'),(16,'羊肉卷','优','个',1.40,0,'yangroujuan'),(17,'香菜','优','斤',15.30,0,'xiangcai'),(18,'大饼','优','个',12.30,0,'dabing'),(19,'龙椒','优','斤',1.04,0,'longjiao'),(20,'枣糕','优','斤',2.04,0,'zaogao'),(21,'面包','优','斤',3.04,0,'mianbao'),(22,'面条','优','斤',1.02,0,'miantiao'),(23,'白菜','优','斤',0.23,0,'baicai');
/*!40000 ALTER TABLE `goods` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-13 15:40:16
