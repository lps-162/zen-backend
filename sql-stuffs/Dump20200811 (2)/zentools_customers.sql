CREATE DATABASE  IF NOT EXISTS `zentools` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `zentools`;
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
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `customerNumber` int NOT NULL,
  `customerName` varchar(50) NOT NULL,
  `phone` varchar(50) NOT NULL,
  `creditLimit` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`customerNumber`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (103,'Atelier graphique','40.32.2555',21000.00),(112,'Signal Gift Stores','7025551838',71800.00),(114,'Australian Collectors, Co.','03 9520 4555',117300.00),(119,'La Rochelle Gifts','40.67.8555',118200.00),(121,'Baane Mini Imports','07-98 9555',81700.00),(124,'Mini Gifts Distributors Ltd.','4155551450',210500.00),(125,'Havel & Zbyszek Co','(26) 642-7555',0.00),(128,'Blauer See Auto, Co.','+49 69 66 90 2555',59700.00),(129,'Mini Wheels Co.','6505555787',64600.00),(131,'Land of Toys Inc.','2125557818',114900.00),(141,'Euro+ Shopping Channel','(91) 555 94 44',227600.00),(144,'Volvo Model Replicas, Co','0921-12 3555',53100.00),(145,'Danish Wholesale Imports','31 12 3555',83400.00),(146,'Saveley & Henriot, Co.','78.32.5555',123900.00),(148,'Dragon Souveniers, Ltd.','+65 221 7555',103800.00),(151,'Muscle Machine Inc','2125557413',138500.00),(157,'Diecast Classics Inc.','2155551555',100600.00),(161,'Technics Stores Inc.','6505556809',84600.00),(166,'Handji Gifts& Co','+65 224 1555',97900.00),(167,'Herkku Gifts','+47 2267 3215',96800.00),(168,'American Souvenirs Inc','2035557845',0.00),(169,'Porto Imports Co.','(1) 356-5555',0.00),(171,'Daedalus Designs Imports','20.16.1555',82900.00),(172,'La Corne D\'abondance, Co.','(1) 42.34.2555',84300.00),(173,'Cambridge Collectables Co.','6175555555',43400.00),(175,'Gift Depot Inc.','2035552570',84300.00),(177,'Osaka Souveniers Co.','+81 06 6342 5555',81200.00),(181,'Vitachrome Inc.','2125551500',76400.00),(186,'Toys of Finland, Co.','90-224 8555',96500.00),(187,'AV Stores, Co.','(171) 555-1555',136800.00),(189,'Clover Collections, Co.','+353 1862 1555',69400.00),(198,'Auto-Moto Classics Inc.','6175558428',23000.00),(201,'UK Collectables, Ltd.','(171) 555-2282',92700.00),(202,'Canadian Gift Exchange Network','(604) 555-3392',90300.00),(204,'Online Mini Collectables','6175557555',68700.00),(205,'Toys4GrownUps.com','6265557265',90700.00),(206,'Asian Shopping Network, Co','+612 9411 1555',0.00),(209,'Mini Caravy','88.60.1555',53800.00),(211,'King Kong Collectables, Co.','+852 2251 1555',58600.00),(216,'Enaco Distributors','(93) 203 4555',60300.00),(219,'Boards & Toys Co.','3105552373',11000.00),(223,'Natürlich Autos','0372-555188',0.00),(227,'Heintze Collectables','86 21 3555',120800.00),(233,'Québec Home Shopping Network','(514) 555-8054',48700.00),(237,'ANG Resellers','(91) 745 6555',0.00),(239,'Collectable Mini Designs Co.','7605558146',105000.00),(240,'giftsbymail.co.uk','(198) 555-8888',93900.00),(242,'Alpha Cognac','61.77.6555',61100.00),(247,'Messner Shopping Network','069-0555984',0.00),(249,'Amica Models & Co.','011-4988555',113000.00),(250,'Lyon Souveniers','+33 1 46 62 7555',68100.00),(256,'Auto Associés & Cie.','30.59.8555',77900.00),(259,'Toms Spezialitäten, Ltd','0221-5554327',120400.00),(260,'Royal Canadian Collectables, Ltd.','(604) 555-4555',89600.00),(273,'Franken Gifts, Co','089-0877555',0.00),(276,'Anna\'s Decorations, Ltd','02 9936 8555',107800.00),(278,'Rovelli Gifts','035-640555',119600.00),(282,'Souveniers And Things Co.','+61 2 9495 8555',93300.00),(286,'Marta\'s Replicas Co.','6175558555',123700.00),(293,'BG&E Collectables','+41 26 425 50 01',0.00),(298,'Vida Sport, Ltd','0897-034555',141300.00),(299,'Norway Gifts By Mail, Co.','+47 2212 1555',95100.00),(303,'Schuyler Imports','+31 20 491 9555',0.00),(307,'Der Hund Imports','030-0074555',0.00),(311,'Oulu Toy Supplies, Inc.','981-443655',90500.00),(314,'Petit Auto','(02) 5554 67',79900.00),(319,'Mini Classics','9145554562',102700.00),(320,'Mini Creations Ltd.','5085559555',94500.00),(321,'Corporate Gift Ideas Co.','6505551386',105000.00),(323,'Down Under Souveniers, Inc','+64 9 312 5555',88000.00),(324,'Stylish Desk Decors, Co.','(171) 555-0297',77000.00),(328,'Tekni Collectables Inc.','2015559350',43000.00),(333,'Australian Gift Network, Co','61-7-3844-6555',51600.00),(334,'Suominen Souveniers','+358 9 8045 555',98800.00),(335,'Cramer Spezialitäten, Ltd','0555-09555',0.00),(339,'Classic Gift Ideas, Inc','2155554695',81100.00),(344,'CAF Imports','+34 913 728 555',59600.00),(347,'Men \'R\' US Retailers, Ltd.','2155554369',57700.00),(348,'Asian Treasures, Inc.','2967 555',0.00),(350,'Marseille Mini Autos','91.24.4555',65000.00),(353,'Reims Collectables','26.47.1555',81100.00),(356,'SAR Distributors, Co','+27 21 550 3555',0.00),(357,'GiftsForHim.com','64-9-3763555',77700.00),(361,'Kommission Auto','0251-555259',0.00),(362,'Gifts4AllAges.com','6175559555',41900.00),(363,'Online Diecast Creations Co.','6035558647',114200.00),(369,'Lisboa Souveniers, Inc','(1) 354-2555',0.00),(376,'Precious Collectables','0452-076555',0.00),(379,'Collectables For Less Inc.','6175558555',70700.00),(381,'Royale Belge','(071) 23 67 2555',23500.00),(382,'Salzburg Collectables','6562-9555',71700.00),(385,'Cruz & Sons Co.','+63 2 555 3587',81500.00),(386,'L\'ordine Souveniers','0522-556555',121400.00),(398,'Tokyo Collectables, Ltd','+81 3 3584 0555',94400.00),(406,'Auto Canal+ Petit','(1) 47.55.6555',95000.00),(409,'Stuttgart Collectable Exchange','0711-555361',0.00),(412,'Extreme Desk Decorations, Ltd','04 499 9555',86800.00),(415,'Bavarian Collectables Imports, Co.',' +49 89 61 08 9555',77000.00),(424,'Classic Legends Inc.','2125558493',67500.00),(443,'Feuer Online Stores, Inc','0342-555176',0.00),(447,'Gift Ideas Corp.','2035554407',49700.00),(448,'Scandinavian Gift Ideas','0695-34 6555',116400.00),(450,'The Sharp Gifts Warehouse','4085553659',77600.00),(452,'Mini Auto Werke','7675-3555',45300.00),(455,'Super Scale Inc.','2035559545',95400.00),(456,'Microscale Inc.','2125551957',39800.00),(458,'Corrida Auto Replicas, Ltd','(91) 555 22 82',104600.00),(459,'Warburg Exchange','0241-039123',0.00),(462,'FunGiftIdeas.com','5085552555',85800.00),(465,'Anton Designs, Ltd.','+34 913 728555',0.00),(471,'Australian Collectables, Ltd','61-9-3844-6555',60300.00),(473,'Frau da Collezione','+39 022515555',34800.00),(475,'West Coast Collectables Co.','3105553722',55400.00),(477,'Mit Vergnügen & Co.','0621-08555',0.00),(480,'Kremlin Collectables, Co.','+7 812 293 0521',0.00),(481,'Raanan Stores, Inc','+ 972 9 959 8555',0.00),(484,'Iberia Gift Imports, Corp.','(95) 555 82 82',65700.00),(486,'Motor Mint Distributors Inc.','2155559857',72600.00),(487,'Signal Collectibles Ltd.','4155554312',60300.00),(489,'Double Decker Gift Stores, Ltd','(171) 555-7555',43300.00),(495,'Diecast Collectables','6175552555',85100.00),(496,'Kelly\'s Gift Shop','+64 9 5555500',110000.00);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-11 11:42:07
