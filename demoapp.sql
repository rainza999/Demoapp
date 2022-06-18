-- MySQL dump 10.13  Distrib 5.7.21, for Win64 (x86_64)
--
-- Host: localhost    Database: demoapp
-- ------------------------------------------------------
-- Server version	5.7.21-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `contacts`
--

DROP TABLE IF EXISTS `contacts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contacts` (
  `id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(225) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(225) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subject` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  `create_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contacts`
--

LOCK TABLES `contacts` WRITE;
/*!40000 ALTER TABLE `contacts` DISABLE KEYS */;
INSERT INTO `contacts` VALUES ('07b47790-c112-4a0e-b372-25ffaeb24548','jann','jannbms@gmail.com','test','test','2022-06-13 23:18:45'),('990db710-a287-4417-a8b9-ecd6c033d80f','namin','navaminsawasdee@gmail.com','test','ทดสอบการส่งเมล์','2022-06-13 23:20:28'),('fa708744-2327-46db-bf76-f737e7a75f0e','jann','jannbms@gmail.com','test','test','2022-06-13 22:29:28');
/*!40000 ALTER TABLE `contacts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `user_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL,
  `useflag` char(1) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_by` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` datetime NOT NULL,
  `update_by` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `update_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('105c29f3-e3bc-4acd-8ba2-fc0e770a9fa3','navamins1','$2a$10$t0XI83WlZCPBIbxpDSXBruJnUgLflpE8FTDNSPwEoAFGIwckG6HFe','Navamin Sawasdee','Y','105c29f3-e3bc-4acd-8ba2-fc0e770a9fa3','2022-06-12 12:45:34','105c29f3-e3bc-4acd-8ba2-fc0e770a9fa3','2022-06-12 12:45:34'),('7dc630c5-f0fd-4230-9f99-5b959496c8e3','navamins','$2a$10$ptsgHaPKJLi0aUm.A2Q8nekQypdypOM.8YYsEXJnMWB4OzH1ZieL.','Navamin Sawasdee','Y','7dc630c5-f0fd-4230-9f99-5b959496c8e3','2022-06-11 16:47:01','7dc630c5-f0fd-4230-9f99-5b959496c8e3','2022-06-12 12:48:07'),('af22e466-ae89-40c9-808a-77a9146c23a7','navamins3','$2a$10$Y5m0MsBvW.IKyT18Q2DZveXLeqQBvNmGAWSxokUQDpAjplZ9TMr5u','Navamin Sawasdee','Y','af22e466-ae89-40c9-808a-77a9146c23a7','2022-06-12 12:46:30','af22e466-ae89-40c9-808a-77a9146c23a7','2022-06-12 12:46:30'),('ce795d00-fb14-43d5-ac91-8dbc219974ed','navamins2','$2a$10$hV8./VGvFB8bY3Tw0cOlSOk6G85vABOBxw6YXETO.ML4gPyaP6DzO','Navamin Sawasdee','Y','ce795d00-fb14-43d5-ac91-8dbc219974ed','2022-06-12 12:50:32','ce795d00-fb14-43d5-ac91-8dbc219974ed','2022-06-12 12:50:32');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-14 21:26:13
