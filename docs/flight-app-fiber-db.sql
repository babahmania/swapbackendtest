-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: 192.168.1.212    Database: flight-app-fiber
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `aircrafts`
--

DROP TABLE IF EXISTS `aircrafts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `aircrafts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `is_active` varchar(1) NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_aircraft_active_status` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `aircrafts`
--

LOCK TABLES `aircrafts` WRITE;
/*!40000 ALTER TABLE `aircrafts` DISABLE KEYS */;
INSERT INTO `aircrafts` VALUES (1,'Boeing 737','1','2021-11-06 04:59:13','2021-11-06 04:59:19',NULL),(2,'Airbus A320','1','2021-11-06 04:59:13','2021-11-06 04:59:19',NULL),(3,'Airbus A330','1','2021-11-06 04:59:13','2021-11-06 04:59:19',NULL),(4,'ATR-72','1','2021-11-06 04:59:13','2021-11-06 04:59:19',NULL);
/*!40000 ALTER TABLE `aircrafts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `airlines`
--

DROP TABLE IF EXISTS `airlines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airlines` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `image_name` varchar(255) NOT NULL,
  `is_active` varchar(1) NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_airline_active_status` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airlines`
--

LOCK TABLES `airlines` WRITE;
/*!40000 ALTER TABLE `airlines` DISABLE KEYS */;
INSERT INTO `airlines` VALUES (1,'Air Asia','airasia.png','1','2021-11-06 04:51:04','2021-11-06 04:51:04',NULL),(2,'Batik Air','batikair.png','1','2021-11-06 04:51:34','2021-11-06 04:51:34',NULL),(3,'Citilink Indonesia','citilink.png','1','2021-11-06 04:51:56','2021-11-06 04:51:56',NULL),(4,'Garuda Indonesia','garuda.png','1','2021-11-06 04:52:11','2021-11-06 04:52:11',NULL),(5,'Lion Air','lionair.png','1','2021-11-06 04:52:24','2021-11-06 04:52:24',NULL);
/*!40000 ALTER TABLE `airlines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `airports`
--

DROP TABLE IF EXISTS `airports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airports` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `code` varchar(4) NOT NULL,
  `city` varchar(100) NOT NULL,
  `country` varchar(100) NOT NULL,
  `is_active` varchar(1) NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`),
  KEY `idx_airport_city` (`city`),
  KEY `idx_airport_active_status` (`is_active`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airports`
--

LOCK TABLES `airports` WRITE;
/*!40000 ALTER TABLE `airports` DISABLE KEYS */;
INSERT INTO `airports` VALUES (1,'Jakarta Indonesia CGK Soekarno Hatta International Airport','CGK','Jakarta','Indonesia','1','2021-11-06 05:02:44','2021-11-06 05:02:44',NULL),(2,'Jakarta Indonesia HLP Halim Perdanakusuma International Airport','HLP','Jakarta','Indonesia','1','2021-11-06 05:03:09','2021-11-06 05:03:09',NULL),(3,'Bali Denpasar Indonesia DPS Ngurah Rai International Airport','DPS','Bali','Indonesia','1','2021-11-06 05:05:45','2021-11-06 05:05:45',NULL),(4,'Yogyakarta Indonesia YKIA International Airport','YKIA','Yogyakarta','Indonesia','1','2021-11-06 05:07:38','2021-11-06 05:07:38',NULL),(5,'Yogyakarta Indonesia JOG Adisucipto','JOG','Yogyakarta','Indonesia','1','2021-11-06 05:08:07','2021-11-06 05:08:07',NULL),(6,'Surabaya Indonesia SUB Juanda','SUB','Surabaya','Indonesia','1','2021-11-06 05:10:27','2021-11-06 05:10:27',NULL);
/*!40000 ALTER TABLE `airports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `flights`
--

DROP TABLE IF EXISTS `flights`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `flights` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `flight_number` varchar(15) NOT NULL,
  `airline_id` int unsigned NOT NULL,
  `origin_id` int unsigned NOT NULL,
  `destination_id` int unsigned NOT NULL,
  `depart_datetime` datetime NOT NULL,
  `arrival_datetime` datetime NOT NULL,
  `duration` varchar(25) NOT NULL,
  `price` int unsigned NOT NULL,
  `seats_available` int unsigned NOT NULL,
  `qty_transit` int unsigned NOT NULL DEFAULT '0',
  `flight_status` varchar(25) NOT NULL,
  `user_id_submit` bigint unsigned NOT NULL,
  `user_id_update` bigint unsigned DEFAULT NULL,
  `user_id_delete` bigint unsigned DEFAULT NULL,
  `transit_first` varchar(100) DEFAULT NULL,
  `transit_second` varchar(100) DEFAULT NULL,
  `transit_third` varchar(100) DEFAULT NULL,
  `is_economy` varchar(1) NOT NULL DEFAULT '1',
  `seats_available_economy` int unsigned NOT NULL DEFAULT '0',
  `is_premium_economy` varchar(1) NOT NULL DEFAULT '1',
  `seats_available_premium_economy` int unsigned NOT NULL DEFAULT '0',
  `is_business` varchar(1) NOT NULL DEFAULT '1',
  `seats_available_business` int unsigned NOT NULL DEFAULT '0',
  `is_first_class` varchar(1) NOT NULL DEFAULT '0',
  `seats_available_first_class` int unsigned NOT NULL DEFAULT '0',
  `qty_baggage` int unsigned NOT NULL DEFAULT '0',
  `qty_cabin` int unsigned NOT NULL DEFAULT '0',
  `aircraft_id` int unsigned NOT NULL,
  `is_meal` varchar(1) NOT NULL DEFAULT '0',
  `is_entertainment` varchar(1) NOT NULL DEFAULT '0',
  `is_power_usb` varchar(1) NOT NULL DEFAULT '0',
  `is_active` varchar(1) NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `flight_number` (`flight_number`),
  KEY `idx_flight_qty_transit` (`qty_transit`),
  KEY `idx_flight_status` (`flight_status`),
  KEY `idx_flight_active_status` (`is_active`),
  KEY `idx_flight_origin_id` (`origin_id`),
  KEY `idx_flight_destination_id` (`destination_id`),
  KEY `idx_flight_user_id_submit` (`user_id_submit`),
  KEY `idx_flight_user_id_update` (`user_id_update`),
  KEY `idx_flight_user_id_delete` (`user_id_delete`),
  KEY `idx_flight_airline_id` (`airline_id`),
  KEY `flights_aircraft_id_aircrafts_id_foreign` (`aircraft_id`),
  CONSTRAINT `flights_aircraft_id_aircrafts_id_foreign` FOREIGN KEY (`aircraft_id`) REFERENCES `aircrafts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `flights_airline_id_airlines_id_foreign` FOREIGN KEY (`airline_id`) REFERENCES `airlines` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `flights_destination_id_airports_id_foreign` FOREIGN KEY (`destination_id`) REFERENCES `airports` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `flights_origin_id_airports_id_foreign` FOREIGN KEY (`origin_id`) REFERENCES `airports` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `flights_user_id_submit_users_id_foreign` FOREIGN KEY (`user_id_submit`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `flights`
--

LOCK TABLES `flights` WRITE;
/*!40000 ALTER TABLE `flights` DISABLE KEYS */;
INSERT INTO `flights` VALUES (1,'CKG-001',5,1,3,'2021-11-06 08:30:00','2021-11-06 09:20:00','1h 50m',1242400,100,0,'open booking',1,0,0,'','','','1',80,'1',15,'1',5,'0',0,20,7,1,'0','1','1','1','2021-11-06 01:29:21','2021-11-06 01:29:21',NULL),(2,'CKG-002',5,1,3,'2021-11-06 10:30:00','2021-11-06 11:20:00','1h 50m',1308400,120,0,'open booking',1,0,0,'','','','1',100,'1',15,'1',5,'0',0,20,7,1,'0','1','1','1','2021-11-06 01:51:19','2021-11-06 01:51:19',NULL),(3,'CKG-003',5,1,3,'2021-11-06 04:00:00','2021-11-06 11:05:00','6h 5m',938700,120,1,'open booking',1,0,0,'Surabaya (SUB) Juanda','','','1',100,'1',15,'1',5,'0',0,20,7,1,'0','1','1','1','2021-11-06 02:33:37','2021-11-06 02:33:37',NULL),(4,'CKG-004',4,1,3,'2021-11-06 07:05:00','2021-11-06 10:00:00','1h 55m',1709100,150,1,'open booking',1,0,0,'Surabaya (SUB) Juanda','','','1',120,'1',10,'1',20,'0',0,20,7,3,'0','1','1','1','2021-11-06 02:36:40','2021-11-06 02:36:40',NULL),(5,'CKG-005',4,1,3,'2021-11-06 16:35:00','2021-11-06 19:30:00','1h 55m',1709100,150,0,'open booking',1,0,0,'','','','1',120,'1',10,'1',20,'0',0,20,7,3,'0','1','1','1','2021-11-06 02:37:59','2021-11-06 02:37:59',NULL),(7,'CKG-006',4,1,3,'2021-11-06 20:35:00','2021-11-06 21:30:00','1h 55m',1709100,150,0,'open booking',1,0,0,'','','','1',120,'1',10,'1',20,'0',0,20,7,3,'0','1','1','1','2021-11-06 06:00:18','2021-11-06 06:00:18',NULL);
/*!40000 ALTER TABLE `flights` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `user_status` varchar(1) NOT NULL DEFAULT '1',
  `user_role` varchar(1) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'babah mania','babahmania@gmail.com','$2a$10$H0FuTrT9mjCATZsEvs.ZnexSHsQicbSdvnGxwaPK7BRA4zqN.7Ioa','1','1','2021-11-06 04:08:11','2021-11-06 04:08:11',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `view_flights`
--

DROP TABLE IF EXISTS `view_flights`;
/*!50001 DROP VIEW IF EXISTS `view_flights`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `view_flights` AS SELECT 
 1 AS `id`,
 1 AS `flight_number`,
 1 AS `airline_id`,
 1 AS `airline_name`,
 1 AS `airline_image_name`,
 1 AS `origin_id`,
 1 AS `origin_name`,
 1 AS `origin_code`,
 1 AS `destination_id`,
 1 AS `destination_name`,
 1 AS `destination_code`,
 1 AS `aircraft_id`,
 1 AS `aircraft_name`,
 1 AS `depart_datetime`,
 1 AS `arrival_datetime`,
 1 AS `duration`,
 1 AS `price`,
 1 AS `seats_available`,
 1 AS `qty_transit`,
 1 AS `flight_status`,
 1 AS `user_id_submit`,
 1 AS `user_id_update`,
 1 AS `user_id_delete`,
 1 AS `transit_first`,
 1 AS `transit_second`,
 1 AS `transit_third`,
 1 AS `is_economy`,
 1 AS `seats_available_economy`,
 1 AS `is_premium_economy`,
 1 AS `seats_available_premium_economy`,
 1 AS `is_business`,
 1 AS `seats_available_business`,
 1 AS `is_first_class`,
 1 AS `seats_available_first_class`,
 1 AS `qty_baggage`,
 1 AS `qty_cabin`,
 1 AS `is_meal`,
 1 AS `is_entertainment`,
 1 AS `is_power_usb`,
 1 AS `is_active`*/;
SET character_set_client = @saved_cs_client;

--
-- Dumping events for database 'flight-app-fiber'
--

--
-- Dumping routines for database 'flight-app-fiber'
--

--
-- Final view structure for view `view_flights`
--

/*!50001 DROP VIEW IF EXISTS `view_flights`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`swap`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `view_flights` AS select `flights`.`id` AS `id`,`flights`.`flight_number` AS `flight_number`,`flights`.`airline_id` AS `airline_id`,`airlines`.`name` AS `airline_name`,`airlines`.`image_name` AS `airline_image_name`,`flights`.`origin_id` AS `origin_id`,`origin_airport`.`name` AS `origin_name`,`origin_airport`.`code` AS `origin_code`,`flights`.`destination_id` AS `destination_id`,`destination_airport`.`name` AS `destination_name`,`destination_airport`.`code` AS `destination_code`,`flights`.`aircraft_id` AS `aircraft_id`,`aircrafts`.`name` AS `aircraft_name`,`flights`.`depart_datetime` AS `depart_datetime`,`flights`.`arrival_datetime` AS `arrival_datetime`,`flights`.`duration` AS `duration`,`flights`.`price` AS `price`,`flights`.`seats_available` AS `seats_available`,`flights`.`qty_transit` AS `qty_transit`,`flights`.`flight_status` AS `flight_status`,`flights`.`user_id_submit` AS `user_id_submit`,`flights`.`user_id_update` AS `user_id_update`,`flights`.`user_id_delete` AS `user_id_delete`,`flights`.`transit_first` AS `transit_first`,`flights`.`transit_second` AS `transit_second`,`flights`.`transit_third` AS `transit_third`,`flights`.`is_economy` AS `is_economy`,`flights`.`seats_available_economy` AS `seats_available_economy`,`flights`.`is_premium_economy` AS `is_premium_economy`,`flights`.`seats_available_premium_economy` AS `seats_available_premium_economy`,`flights`.`is_business` AS `is_business`,`flights`.`seats_available_business` AS `seats_available_business`,`flights`.`is_first_class` AS `is_first_class`,`flights`.`seats_available_first_class` AS `seats_available_first_class`,`flights`.`qty_baggage` AS `qty_baggage`,`flights`.`qty_cabin` AS `qty_cabin`,`flights`.`is_meal` AS `is_meal`,`flights`.`is_entertainment` AS `is_entertainment`,`flights`.`is_power_usb` AS `is_power_usb`,`flights`.`is_active` AS `is_active` from ((((`flights` join `airlines`) join `airports` `origin_airport`) join `airports` `destination_airport`) join `aircrafts`) where ((`flights`.`airline_id` = `airlines`.`id`) and (`flights`.`origin_id` = `origin_airport`.`id`) and (`flights`.`destination_id` = `destination_airport`.`id`) and (`flights`.`aircraft_id` = `aircrafts`.`id`)) order by `flights`.`depart_datetime` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-07 23:47:14

GRANT ALL PRIVILEGES ON *.* TO 'swap'@'%' WITH GRANT OPTION;