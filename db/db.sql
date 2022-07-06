-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: mvc_go
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Table structure for table `addresses`
--

DROP TABLE IF EXISTS `addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `addresses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `neighborhood` varchar(250) NOT NULL,
  `street` varchar(150) NOT NULL,
  `number` int NOT NULL,
  `id_user` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
INSERT INTO `addresses` VALUES (1,'Villa Cabrera','Curuzu Cuatia',2606,1),(2,'La reserva','El cuñacaico',9,2),(3,'Villa Allende','No se',2550,3);
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Celulares','Dispositivos Moviles'),(2,'Camaras','Camaras Digitales'),(3,'Consolas','Consolas'),(4,'Computacion','Componentes de PC, PC'),(5,'Drones','Drones a control remoto'),(6,'Televisores','Pantallas de distinto tamaño');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `precio_unitario` decimal(60,4) NOT NULL,
  `cantidad` decimal(60,4) NOT NULL,
  `total` decimal(60,4) NOT NULL,
  `id_product` int NOT NULL,
  `id_order` int NOT NULL,
  `nombre` varchar(550) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_details`
--

LOCK TABLES `order_details` WRITE;
/*!40000 ALTER TABLE `order_details` DISABLE KEYS */;
INSERT INTO `order_details` VALUES (5,1500.0000,1.0000,1500.0000,5,5,'Canon EOS T100'),(6,1200.0000,4.0000,4800.0000,8,5,'Sony Playstation 5'),(7,700.0000,1.0000,700.0000,2,6,'Iphone 11'),(8,600.0000,1.0000,600.0000,1,6,'Iphone 10'),(9,700.0000,1.0000,700.0000,2,7,'Iphone 11'),(10,600.0000,1.0000,600.0000,1,7,'Iphone 10'),(11,1500.0000,5.0000,7500.0000,5,8,'Canon EOS T100'),(12,1200.0000,10.0000,12000.0000,8,8,'Sony Playstation 5'),(13,700.0000,1.0000,700.0000,2,9,'Iphone 11'),(14,600.0000,1.0000,600.0000,1,9,'Iphone 10'),(15,700.0000,1.0000,700.0000,2,10,'Iphone 11'),(16,600.0000,20.0000,12000.0000,1,10,'Iphone 10'),(17,700.0000,1.0000,700.0000,2,11,'Iphone 11'),(18,600.0000,1.0000,600.0000,1,11,'Iphone 10'),(19,700.0000,1.0000,700.0000,2,12,'Iphone 11'),(20,600.0000,1.0000,600.0000,1,12,'Iphone 10'),(21,1200.0000,1.0000,1200.0000,7,13,'Xbox Series X'),(22,120.0000,1.0000,120.0000,15,13,'Drone Gadnic'),(23,1200.0000,2.0000,2400.0000,7,15,'Xbox Series X'),(24,1700.0000,2.0000,3400.0000,10,16,'Nvidia GeForce 1060');
/*!40000 ALTER TABLE `order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monto_final` decimal(60,4) NOT NULL,
  `fecha` datetime NOT NULL,
  `id_user` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (5,6300.0000,'2022-06-09 04:53:39',3),(6,1300.0000,'2022-06-09 05:04:20',3),(7,1300.0000,'2022-06-09 05:06:08',3),(8,19500.0000,'2022-06-14 15:28:01',3),(9,1300.0000,'2022-06-14 20:01:16',3),(10,12700.0000,'2022-06-14 20:02:20',3),(11,1300.0000,'2022-06-30 14:32:30',3),(12,1300.0000,'2022-06-30 15:05:21',2),(13,1320.0000,'2022-07-04 14:03:12',1),(15,2400.0000,'2022-07-04 17:32:34',3),(16,3400.0000,'2022-07-04 17:33:52',1);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(350) NOT NULL,
  `price` decimal(60,4) NOT NULL,
  `stock` int NOT NULL,
  `id_category` int NOT NULL,
  `description` varchar(350) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Iphone 10',600.0000,9,1,'128gb Azul'),(2,'Iphone 11',700.0000,8,1,'256gb Verde'),(3,'Iphone 12',800.0000,10,1,'516gb Negro'),(4,'Nikon Coolpix B500',1200.0000,5,2,'Compacta Avanzada color negro'),(5,'Canon EOS T100',1500.0000,4,2,'55mm DSLR negro'),(6,'Canon EOS M200',1000.0000,3,2,'Sin espejo'),(7,'Xbox Series X',1200.0000,7,3,'1TB color negro'),(8,'Sony Playstation 5',1200.0000,10,3,'Standard'),(9,'Nintendo Switch',800.0000,4,3,'32gb standard'),(10,'Nvidia GeForce 1060',1700.0000,0,4,'8gb Gigabyte'),(11,'Intel I3 10100',900.0000,20,4,'4 nucleos'),(12,'AMD Ryzen 7 5700G',4.0000,2000,5,'8 nucleos'),(13,'Mini drone DJI Mavic',800.0000,5,5,'Light gray 3 baterias'),(14,'Drone Samsung A9',400.0000,1,5,'No existe, lo invente'),(15,'Drone Gadnic',120.0000,0,5,'Camara HD'),(16,'Smart TV Samsung',500.0000,5,6,'LED 4K'),(17,'LG Smart TV',400.0000,2,6,'4K 55pulgadas'),(18,'Smart TV Noblex',200.0000,100,6,'HD 32pulgadas');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(350) NOT NULL,
  `last_name` varchar(250) NOT NULL,
  `user_name` varchar(150) NOT NULL,
  `password` varchar(150) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Matias','Lessio','matiaslessio','mati123'),(2,'Octavio','Garcia','pollo','pollo123'),(3,'Benjamin','Angelone','benja','benja123');
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

-- Dump completed on 2022-07-06 14:19:28
