-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: zentools
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `orderNumber` int NOT NULL,
  `orderDate` date NOT NULL,
  `shippedDate` date DEFAULT NULL,
  `status` varchar(15) NOT NULL,
  `customerNumber` int NOT NULL,
  PRIMARY KEY (`orderNumber`),
  KEY `orders_cust_fk` (`customerNumber`),
  CONSTRAINT `orders_cust_fk` FOREIGN KEY (`customerNumber`) REFERENCES `customers` (`customerNumber`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (10100,'2003-01-06','2003-01-10','Shipped',363),(10101,'2003-01-09','2003-01-11','Shipped',128),(10102,'2003-01-10','2003-01-14','Shipped',181),(10103,'2003-01-29','2003-02-02','Shipped',121),(10104,'2003-01-31','2003-02-01','Shipped',141),(10105,'2003-02-11','2003-02-12','Shipped',145),(10106,'2003-02-17','2003-02-21','Shipped',278),(10107,'2003-02-24','2003-02-26','Shipped',131),(10108,'2003-03-03','2003-03-08','Shipped',385),(10109,'2003-03-10','2003-03-11','Shipped',486),(10110,'2003-03-18','2003-03-20','Shipped',187),(10111,'2003-03-25','2003-03-30','Shipped',129),(10112,'2003-03-24','2003-03-29','Shipped',144),(10113,'2003-03-26','2003-03-27','Shipped',124),(10114,'2003-04-01','2003-04-02','Shipped',172),(10115,'2003-04-04','2003-04-07','Shipped',424),(10116,'2003-04-11','2003-04-13','Shipped',381),(10117,'2003-04-16','2003-04-17','Shipped',148),(10118,'2003-04-21','2003-04-26','Shipped',216),(10119,'2003-04-28','2003-05-02','Shipped',382),(10120,'2003-04-29','2003-05-01','Shipped',114),(10121,'2003-05-07','2003-05-13','Shipped',353),(10122,'2003-05-08','2003-05-13','Shipped',350),(10123,'2003-05-20','2003-05-22','Shipped',103),(10124,'2003-05-21','2003-05-25','Shipped',112),(10125,'2003-05-21','2003-05-24','Shipped',114),(10126,'2003-05-28','2003-06-02','Shipped',458),(10127,'2003-06-03','2003-06-06','Shipped',151),(10128,'2003-06-06','2003-06-11','Shipped',141),(10129,'2003-06-12','2003-06-14','Shipped',324),(10130,'2003-06-16','2003-06-21','Shipped',198),(10131,'2003-06-16','2003-06-21','Shipped',447),(10132,'2003-06-25','2003-06-28','Shipped',323),(10133,'2003-06-27','2003-07-03','Shipped',141),(10134,'2003-07-01','2003-07-05','Shipped',250),(10135,'2003-07-02','2003-07-03','Shipped',124),(10136,'2003-07-04','2003-07-06','Shipped',242),(10137,'2003-07-10','2003-07-14','Shipped',353),(10138,'2003-07-07','2003-07-13','Shipped',496),(10139,'2003-07-16','2003-07-21','Shipped',282),(10140,'2003-07-24','2003-07-30','Shipped',161),(10141,'2003-08-01','2003-08-04','Shipped',334),(10142,'2003-08-08','2003-08-13','Shipped',124),(10143,'2003-08-10','2003-08-12','Shipped',320),(10144,'2003-08-13','2003-08-14','Shipped',381),(10145,'2003-08-25','2003-08-31','Shipped',205),(10146,'2003-09-03','2003-09-06','Shipped',447),(10147,'2003-09-05','2003-09-09','Shipped',379),(10148,'2003-09-11','2003-09-15','Shipped',276),(10149,'2003-09-12','2003-09-17','Shipped',487),(10150,'2003-09-19','2003-09-21','Shipped',148),(10151,'2003-09-21','2003-09-24','Shipped',311),(10152,'2003-09-25','2003-10-01','Shipped',333),(10153,'2003-09-28','2003-10-03','Shipped',141),(10154,'2003-10-02','2003-10-08','Shipped',219),(10155,'2003-10-06','2003-10-07','Shipped',186),(10156,'2003-10-08','2003-10-11','Shipped',141),(10157,'2003-10-09','2003-10-14','Shipped',473),(10158,'2003-10-10','2003-10-15','Shipped',121),(10159,'2003-10-10','2003-10-16','Shipped',321),(10160,'2003-10-11','2003-10-17','Shipped',347),(10161,'2003-10-17','2003-10-20','Shipped',227),(10162,'2003-10-18','2003-10-19','Shipped',321),(10163,'2003-10-20','2003-10-24','Shipped',424),(10164,'2003-10-21','2003-10-23','Resolved',452),(10165,'2003-10-22','2003-12-26','Shipped',148),(10166,'2003-10-21','2003-10-27','Shipped',462),(10167,'2003-10-23',NULL,'Cancelled',448),(10168,'2003-10-28','2003-11-01','Shipped',161),(10169,'2003-11-04','2003-11-09','Shipped',276),(10170,'2003-11-04','2003-11-07','Shipped',452),(10171,'2003-11-05','2003-11-07','Shipped',233),(10172,'2003-11-05','2003-11-11','Shipped',175),(10173,'2003-11-05','2003-11-09','Shipped',278),(10174,'2003-11-06','2003-11-10','Shipped',333),(10175,'2003-11-06','2003-11-09','Shipped',324),(10176,'2003-11-06','2003-11-12','Shipped',386),(10177,'2003-11-07','2003-11-12','Shipped',344),(10178,'2003-11-08','2003-11-10','Shipped',242),(10179,'2003-11-11','2003-11-13','Cancelled',496),(10180,'2003-11-11','2003-11-14','Shipped',171),(10181,'2003-11-12','2003-11-15','Shipped',167),(10182,'2003-11-12','2003-11-18','Shipped',124),(10183,'2003-11-13','2003-11-15','Shipped',339),(10184,'2003-11-14','2003-11-20','Shipped',484),(10185,'2003-11-14','2003-11-20','Shipped',320),(10186,'2003-11-14','2003-11-18','Shipped',489),(10187,'2003-11-15','2003-11-16','Shipped',211),(10188,'2003-11-18','2003-11-24','Shipped',167),(10189,'2003-11-18','2003-11-24','Shipped',205),(10190,'2003-11-19','2003-11-20','Shipped',141),(10191,'2003-11-20','2003-11-24','Shipped',259),(10192,'2003-11-20','2003-11-25','Shipped',363),(10193,'2003-11-21','2003-11-27','Shipped',471),(10194,'2003-11-25','2003-11-26','Shipped',146),(10195,'2003-11-25','2003-11-28','Shipped',319),(10196,'2003-11-26','2003-12-01','Shipped',455),(10197,'2003-11-26','2003-12-01','Shipped',216),(10198,'2003-11-27','2003-12-03','Shipped',385),(10199,'2003-12-01','2003-12-06','Shipped',475),(10200,'2003-12-01','2003-12-06','Shipped',211),(10201,'2003-12-01','2003-12-02','Shipped',129),(10202,'2003-12-02','2003-12-06','Shipped',357),(10203,'2003-12-02','2003-12-07','Shipped',141),(10204,'2003-12-02','2003-12-04','Shipped',151),(10205,'2003-12-03','2003-12-07','Shipped',141),(10206,'2003-12-05','2003-12-08','Shipped',202),(10207,'2003-12-09','2003-12-11','Shipped',495),(10208,'2004-01-02','2004-01-04','Shipped',146),(10209,'2004-01-09','2004-01-12','Shipped',347),(10210,'2004-01-12','2004-01-20','Shipped',177),(10211,'2004-01-15','2004-01-18','Shipped',406),(10212,'2004-01-16','2004-01-18','Shipped',141),(10213,'2004-01-22','2004-01-27','Shipped',489),(10214,'2004-01-26','2004-01-29','Shipped',458),(10215,'2004-01-29','2004-02-01','Shipped',475),(10216,'2004-02-02','2004-02-04','Shipped',256),(10217,'2004-02-04','2004-02-06','Shipped',166),(10218,'2004-02-09','2004-02-11','Shipped',473),(10219,'2004-02-10','2004-02-12','Shipped',487),(10220,'2004-02-12','2004-02-16','Shipped',189),(10221,'2004-02-18','2004-02-19','Shipped',314),(10222,'2004-02-19','2004-02-20','Shipped',239),(10223,'2004-02-20','2004-02-24','Shipped',114),(10224,'2004-02-21','2004-02-26','Shipped',171),(10225,'2004-02-22','2004-02-24','Shipped',298),(10226,'2004-02-26','2004-03-02','Shipped',239),(10227,'2004-03-02','2004-03-08','Shipped',146),(10228,'2004-03-10','2004-03-13','Shipped',173),(10229,'2004-03-11','2004-03-12','Shipped',124),(10230,'2004-03-15','2004-03-20','Shipped',128),(10231,'2004-03-19','2004-03-25','Shipped',344),(10232,'2004-03-20','2004-03-25','Shipped',240),(10233,'2004-03-29','2004-04-02','Shipped',328),(10234,'2004-03-30','2004-04-02','Shipped',412),(10235,'2004-04-02','2004-04-06','Shipped',260),(10236,'2004-04-03','2004-04-08','Shipped',486),(10237,'2004-04-05','2004-04-10','Shipped',181),(10238,'2004-04-09','2004-04-10','Shipped',145),(10239,'2004-04-12','2004-04-17','Shipped',311),(10240,'2004-04-13','2004-04-20','Shipped',177),(10241,'2004-04-13','2004-04-19','Shipped',209),(10242,'2004-04-20','2004-04-25','Shipped',456),(10243,'2004-04-26','2004-04-28','Shipped',495),(10244,'2004-04-29','2004-05-04','Shipped',141),(10245,'2004-05-04','2004-05-09','Shipped',455),(10246,'2004-05-05','2004-05-06','Shipped',141),(10247,'2004-05-05','2004-05-08','Shipped',334),(10248,'2004-05-07',NULL,'Cancelled',131),(10249,'2004-05-08','2004-05-11','Shipped',173),(10250,'2004-05-11','2004-05-15','Shipped',450),(10251,'2004-05-18','2004-05-24','Shipped',328),(10252,'2004-05-26','2004-05-29','Shipped',406),(10253,'2004-06-01','2004-06-02','Cancelled',201),(10254,'2004-06-03','2004-06-04','Shipped',323),(10255,'2004-06-04','2004-06-09','Shipped',209),(10256,'2004-06-08','2004-06-10','Shipped',145),(10257,'2004-06-14','2004-06-15','Shipped',450),(10258,'2004-06-15','2004-06-23','Shipped',398),(10259,'2004-06-15','2004-06-17','Shipped',166),(10260,'2004-06-16',NULL,'Cancelled',357),(10261,'2004-06-17','2004-06-22','Shipped',233),(10262,'2004-06-24',NULL,'Cancelled',141),(10263,'2004-06-28','2004-07-02','Shipped',175),(10264,'2004-06-30','2004-07-01','Shipped',362),(10265,'2004-07-02','2004-07-07','Shipped',471),(10266,'2004-07-06','2004-07-10','Shipped',386),(10267,'2004-07-07','2004-07-09','Shipped',151),(10268,'2004-07-12','2004-07-14','Shipped',412),(10269,'2004-07-16','2004-07-18','Shipped',382),(10270,'2004-07-19','2004-07-24','Shipped',282),(10271,'2004-07-20','2004-07-23','Shipped',124),(10272,'2004-07-20','2004-07-22','Shipped',157),(10273,'2004-07-21','2004-07-22','Shipped',314),(10274,'2004-07-21','2004-07-22','Shipped',379),(10275,'2004-07-23','2004-07-29','Shipped',119),(10276,'2004-08-02','2004-08-08','Shipped',204),(10277,'2004-08-04','2004-08-05','Shipped',148),(10278,'2004-08-06','2004-08-09','Shipped',112),(10279,'2004-08-09','2004-08-15','Shipped',141),(10280,'2004-08-17','2004-08-19','Shipped',249),(10281,'2004-08-19','2004-08-23','Shipped',157),(10282,'2004-08-20','2004-08-22','Shipped',124),(10283,'2004-08-20','2004-08-23','Shipped',260),(10284,'2004-08-21','2004-08-26','Shipped',299),(10285,'2004-08-27','2004-08-31','Shipped',286),(10286,'2004-08-28','2004-09-01','Shipped',172),(10287,'2004-08-30','2004-09-01','Shipped',298),(10288,'2004-09-01','2004-09-05','Shipped',166),(10289,'2004-09-03','2004-09-04','Shipped',167),(10290,'2004-09-07','2004-09-13','Shipped',198),(10291,'2004-09-08','2004-09-14','Shipped',448),(10292,'2004-09-08','2004-09-11','Shipped',131),(10293,'2004-09-09','2004-09-14','Shipped',249),(10294,'2004-09-10','2004-09-14','Shipped',204),(10295,'2004-09-10','2004-09-14','Shipped',362),(10296,'2004-09-15','2004-09-16','Shipped',415),(10297,'2004-09-16','2004-09-21','Shipped',189),(10298,'2004-09-27','2004-10-01','Shipped',103),(10299,'2004-09-30','2004-10-01','Shipped',186),(10300,'2003-10-04','2003-10-09','Shipped',128),(10301,'2003-10-05','2003-10-08','Shipped',299),(10302,'2003-10-06','2003-10-07','Shipped',201),(10303,'2004-10-06','2004-10-09','Shipped',484),(10304,'2004-10-11','2004-10-17','Shipped',256),(10305,'2004-10-13','2004-10-15','Shipped',286),(10306,'2004-10-14','2004-10-17','Shipped',187),(10307,'2004-10-14','2004-10-20','Shipped',339),(10308,'2004-10-15','2004-10-20','Shipped',319),(10309,'2004-10-15','2004-10-18','Shipped',121),(10310,'2004-10-16','2004-10-18','Shipped',259),(10311,'2004-10-16','2004-10-20','Shipped',141),(10312,'2004-10-21','2004-10-23','Shipped',124),(10313,'2004-10-22','2004-10-25','Shipped',202),(10314,'2004-10-22','2004-10-23','Shipped',227),(10315,'2004-10-29','2004-10-30','Shipped',119),(10316,'2004-11-01','2004-11-07','Shipped',240),(10317,'2004-11-02','2004-11-08','Shipped',161),(10318,'2004-11-02','2004-11-07','Shipped',157),(10319,'2004-11-03','2004-11-06','Shipped',456),(10320,'2004-11-03','2004-11-07','Shipped',144),(10321,'2004-11-04','2004-11-07','Shipped',462),(10322,'2004-11-04','2004-11-10','Shipped',363),(10323,'2004-11-05','2004-11-09','Shipped',128),(10324,'2004-11-05','2004-11-08','Shipped',181),(10325,'2004-11-05','2004-11-08','Shipped',121),(10326,'2004-11-09','2004-11-10','Shipped',144),(10327,'2004-11-10','2004-11-13','Resolved',145),(10328,'2004-11-12','2004-11-18','Shipped',278),(10329,'2004-11-15','2004-11-16','Shipped',131),(10330,'2004-11-16','2004-11-21','Shipped',385),(10331,'2004-11-17','2004-11-23','Shipped',486),(10332,'2004-11-17','2004-11-18','Shipped',187),(10333,'2004-11-18','2004-11-20','Shipped',129),(10334,'2004-11-19',NULL,'On Hold',144),(10335,'2004-11-19','2004-11-23','Shipped',124),(10336,'2004-11-20','2004-11-24','Shipped',172),(10337,'2004-11-21','2004-11-26','Shipped',424),(10338,'2004-11-22','2004-11-27','Shipped',381),(10339,'2004-11-23','2004-11-30','Shipped',398),(10340,'2004-11-24','2004-11-25','Shipped',216),(10341,'2004-11-24','2004-11-29','Shipped',382),(10342,'2004-11-24','2004-11-29','Shipped',114),(10343,'2004-11-24','2004-11-26','Shipped',353),(10344,'2004-11-25','2004-11-29','Shipped',350),(10345,'2004-11-25','2004-11-26','Shipped',103),(10346,'2004-11-29','2004-11-30','Shipped',112),(10347,'2004-11-29','2004-11-30','Shipped',114),(10348,'2004-11-01','2004-11-05','Shipped',458),(10349,'2004-12-01','2004-12-03','Shipped',151),(10350,'2004-12-02','2004-12-05','Shipped',141),(10351,'2004-12-03','2004-12-07','Shipped',324),(10352,'2004-12-03','2004-12-09','Shipped',198),(10353,'2004-12-04','2004-12-05','Shipped',447),(10354,'2004-12-04','2004-12-05','Shipped',323),(10355,'2004-12-07','2004-12-13','Shipped',141),(10356,'2004-12-09','2004-12-12','Shipped',250),(10357,'2004-12-10','2004-12-14','Shipped',124),(10358,'2004-12-10','2004-12-16','Shipped',141),(10359,'2004-12-15','2004-12-18','Shipped',353),(10360,'2004-12-16','2004-12-18','Shipped',496),(10361,'2004-12-17','2004-12-20','Shipped',282),(10362,'2005-01-05','2005-01-10','Shipped',161),(10363,'2005-01-06','2005-01-10','Shipped',334),(10364,'2005-01-06','2005-01-09','Shipped',350),(10365,'2005-01-07','2005-01-11','Shipped',320),(10366,'2005-01-10','2005-01-12','Shipped',381),(10367,'2005-01-12','2005-01-16','Resolved',205),(10368,'2005-01-19','2005-01-24','Shipped',124),(10369,'2005-01-20','2005-01-24','Shipped',379),(10370,'2005-01-20','2005-01-25','Shipped',276),(10371,'2005-01-23','2005-01-25','Shipped',124),(10372,'2005-01-26','2005-01-28','Shipped',398),(10373,'2005-01-31','2005-02-06','Shipped',311),(10374,'2005-02-02','2005-02-03','Shipped',333),(10375,'2005-02-03','2005-02-06','Shipped',119),(10376,'2005-02-08','2005-02-13','Shipped',219),(10377,'2005-02-09','2005-02-12','Shipped',186),(10378,'2005-02-10','2005-02-11','Shipped',141),(10379,'2005-02-10','2005-02-11','Shipped',141),(10380,'2005-02-16','2005-02-18','Shipped',141),(10381,'2005-02-17','2005-02-18','Shipped',321),(10382,'2005-02-17','2005-02-18','Shipped',124),(10383,'2005-02-22','2005-02-25','Shipped',141),(10384,'2005-02-23','2005-02-27','Shipped',321),(10385,'2005-02-28','2005-03-01','Shipped',124),(10386,'2005-03-01','2005-03-06','Resolved',141),(10387,'2005-03-02','2005-03-06','Shipped',148),(10388,'2005-03-03','2005-03-09','Shipped',462),(10389,'2005-03-03','2005-03-08','Shipped',448),(10390,'2005-03-04','2005-03-07','Shipped',124),(10391,'2005-03-09','2005-03-15','Shipped',276),(10392,'2005-03-10','2005-03-12','Shipped',452),(10393,'2005-03-11','2005-03-14','Shipped',323),(10394,'2005-03-15','2005-03-19','Shipped',141),(10395,'2005-03-17','2005-03-23','Shipped',250),(10396,'2005-03-23','2005-03-28','Shipped',124),(10397,'2005-03-28','2005-04-01','Shipped',242),(10398,'2005-03-30','2005-03-31','Shipped',353),(10399,'2005-04-01','2005-04-03','Shipped',496),(10400,'2005-04-01','2005-04-04','Shipped',450),(10401,'2005-04-03',NULL,'On Hold',328),(10402,'2005-04-07','2005-04-12','Shipped',406),(10403,'2005-04-08','2005-04-11','Shipped',201),(10404,'2005-04-08','2005-04-11','Shipped',323),(10405,'2005-04-14','2005-04-20','Shipped',209),(10406,'2005-04-15','2005-04-21','Disputed',145),(10407,'2005-04-22',NULL,'On Hold',450),(10408,'2005-04-22','2005-04-27','Shipped',398),(10409,'2005-04-23','2005-04-24','Shipped',166),(10410,'2005-04-29','2005-04-30','Shipped',357),(10411,'2005-05-01','2005-05-06','Shipped',233),(10412,'2005-05-03','2005-05-05','Shipped',141),(10413,'2005-05-05','2005-05-09','Shipped',175),(10414,'2005-05-06',NULL,'On Hold',362),(10415,'2005-05-09','2005-05-12','Disputed',471),(10416,'2005-05-10','2005-05-14','Shipped',386),(10417,'2005-05-13','2005-05-19','Disputed',141),(10418,'2005-05-16','2005-05-20','Shipped',412),(10419,'2005-05-17','2005-05-19','Shipped',382),(10420,'2005-05-29',NULL,'In Process',282),(10421,'2005-05-29',NULL,'In Process',124),(10422,'2005-05-30',NULL,'In Process',157),(10423,'2005-05-30',NULL,'In Process',314),(10424,'2005-05-31',NULL,'In Process',141),(10425,'2005-05-31',NULL,'In Process',119);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-11 10:24:41