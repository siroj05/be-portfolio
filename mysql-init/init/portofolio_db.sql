-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: portfolio
-- ------------------------------------------------------
-- Server version	8.0.43

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
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES ('2f0d08c0-9c77-41bf-bd1b-9426a6bf5454','Languages'),('d46303c5-91f0-4f7a-92a4-b6771a3209c2','Styling & Design'),('de6a79c3-2379-4e33-8b96-3566f7d7e593','Backend & Databases'),('e6004ede-e6f3-4eba-afab-f7183802d64f','Frontend Frameworks');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `experiences`
--

DROP TABLE IF EXISTS `experiences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `experiences` (
  `id` varchar(36) NOT NULL,
  `office` varchar(255) NOT NULL,
  `position` varchar(255) NOT NULL,
  `start` date NOT NULL,
  `end` date DEFAULT NULL,
  `description` text,
  `present` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `experiences`
--

LOCK TABLES `experiences` WRITE;
/*!40000 ALTER TABLE `experiences` DISABLE KEYS */;
INSERT INTO `experiences` VALUES ('407a822c-d49e-4cba-8d69-580faa02f536','PT. Arthamas Solusindo','Frontend Developer','2023-01-01',NULL,'\"<ul class=\\\"list-node\\\"><li><p class=\\\"text-node\\\">Developed and maintained ERP web applications to streamline business operations and improve efficiency.</p></li><li><p class=\\\"text-node\\\">Built and optimized responsive user interfaces using React.js, TypeScript, and Tailwind CSS.</p></li><li><p class=\\\"text-node\\\">Worked closely with backend developers to integrate RESTful APIs for real-time data processing.</p></li><li><p class=\\\"text-node\\\">Improved application performance and usability, reducing load time and enhancing user experience.</p></li><li><p class=\\\"text-node\\\">Collaborated with cross-functional teams to deliver scalable and maintainable solutions aligned with business requirements.</p></li></ul><p class=\\\"text-node\\\"></p>\"',1),('86682e68-3aca-43ad-b819-fda42550ae7f','Software Engineering and Application (SEA)','Laboratory Assistant','2020-08-01','2021-09-30','\"<ul class=\\\"list-node\\\"><li><p class=\\\"text-node\\\">Conducted C and C++ programming classes, guiding students in fundamental programming concepts and problem-solving techniques.</p></li><li><p class=\\\"text-node\\\">Taught web development, covering HTML, CSS, and JavaScript to help students build interactive websites.</p></li><li><p class=\\\"text-node\\\">Assisted in debugging and troubleshooting students code, ensuring they understood best practices in software development.</p></li><li><p class=\\\"text-node\\\">Provided one-on-one mentoring and practical sessions to reinforce theoretical knowledge with hands-on coding exercises.</p></li><li><p class=\\\"text-node\\\">Collaborated with fellow assistants and lecturers to improve learning materials and teaching methods.</p></li></ul><p class=\\\"text-node\\\"></p>\"',0),('ae068417-0746-4678-adaf-5f911357cb4c','Indonesian Institute of Science (LIPI) - Internship','Frontend Developer','2021-07-01','2021-08-31','\"<ul class=\\\"list-node\\\"><li><p class=\\\"text-node\\\">Built a responsive user interface (UI) using Vue.js and CSS, ensuring accessibility across different devices.</p></li><li><p class=\\\"text-node\\\">Collaborated with the backend team to ensure seamless data synchronization between the roasting machine and the dashboard.</p></li><li><p class=\\\"text-node\\\">Developed an interactive dashboard to monitor an automated coffee roasting machine.</p></li></ul><p class=\\\"text-node\\\"></p>\"',0);
/*!40000 ALTER TABLE `experiences` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `messages`
--

DROP TABLE IF EXISTS `messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `messages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `message` text NOT NULL,
  `is_read` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
/*!40000 ALTER TABLE `messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `profile` (
  `id` varchar(36) NOT NULL,
  `user_id` int NOT NULL,
  `image_path` varchar(512) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `job_title` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `linkedin` varchar(255) DEFAULT NULL,
  `repo` varchar(255) DEFAULT NULL,
  `about` text,
  `phone_number` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_profile` (`user_id`),
  CONSTRAINT `fk_user_profile` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profile`
--

LOCK TABLES `profile` WRITE;
/*!40000 ALTER TABLE `profile` DISABLE KEYS */;
INSERT INTO `profile` VALUES ('626d705c-dc31-49fe-80a8-a418e5279af0',1,'uploads/1758701110711036100_cropped-image.png','Siroojuddin Apendi','Frontend Developer','rojudin123@gmail.com','https://www.linkedin.com/in/siroojuddin-apendi-121a6a1a0/','https://github.com/siroj05','Passionate Frontend Developer with over 2 years of experience in developing scalable, high-performance web applications. Skilled in React.js, TypeScript, and Tailwind CSS, with a strong emphasis on clean, maintainable code and efficient user interactions. Enthusiastic about learning new technologies, tackling complex challenges, and staying up to date with industry advancements. Highly adaptable and always eager to explore new programming languages and frameworks to build innovative digital solutions.','+6281383304270','Tangerang, Banten');
/*!40000 ALTER TABLE `profile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `projects` (
  `id` varchar(36) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` text,
  `tech_stack` varchar(255) DEFAULT NULL,
  `demo_url` varchar(255) DEFAULT NULL,
  `github_url` varchar(255) DEFAULT NULL,
  `filepath` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES ('599044cc-3285-40b1-9cea-e4679dfce0e0','Roxy Square Company Profile','A company profile website for Roxy Square. Built to showcase their retail tenants, location info, and customer services.','React,Vite, Ant Design. Tailwind CSS','https://roxysquarejakarta.com/','','uploads/compro.png'),('5ab01648-d19c-45e9-815c-315512043c2e','Post Web','A simple blog-like post web built with React that fetches and displays post data from the DummyJSON API. Users can view a list of posts and navigate to individual post details.','Next.js, Tailwind CSS','https://posts-test-site.netlify.app/','https://github.com/siroj05/Posts','uploads/post-web.png'),('902d2f91-4cf6-424e-a592-4e1df9d30f84','Pokedex','A simple Pokédex web application built using Next.js and the PokeAPI. Users can search for Pokémon, view detailed information, and explore different types and stats.','Next.js, Zustand, Tailwind CSS, Shadcn UI, Typescript','https://siroj-pokedex.netlify.app/list-pokemon','https://github.com/siroj05/pokemon','uploads/pokedex.png'),('bb1424de-65d1-4b65-9ac4-565682eb8353','Hoyolab Clone Backend','A hoyolab clone backend.','Express.js, MongoDB','','https://github.com/siroj05/express-mongo','uploads/nodejs240.png'),('c1cb8794-6e9d-4c74-8fb3-320639418e90','Hoyolab Clone','A Hoyolab-inspired clone with user registration, login, posting, and comment features.','React, Tailwind CSS, TypeScript, Redux, Vite','https://siroj-hoyolab-clone.netlify.app/home','https://github.com/siroj05/hoyolab-clone','uploads/hoyolab-clone.png'),('cfc6a6fc-e419-4781-aa0f-7f2583310cf1','ERP Project','A financial ERP system developed for PT Arthamas Solusindo, designed to manage transactions, budgeting, and generate financial reports efficiently.','Remix.js, Zustand, Tailwind CSS, Shadcn UI, Typescript','','','uploads/remix-run.png');
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `skills`
--

DROP TABLE IF EXISTS `skills`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `skills` (
  `id` varchar(36) NOT NULL,
  `category_id` varchar(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `icon` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `skills_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `skills`
--

LOCK TABLES `skills` WRITE;
/*!40000 ALTER TABLE `skills` DISABLE KEYS */;
INSERT INTO `skills` VALUES ('173adb8c-4c6d-4adb-8650-88b4c6b49dc3','d46303c5-91f0-4f7a-92a4-b6771a3209c2','Tailwind CSS','uploads/1758875274_tailwindcss.png'),('1bb1c7da-305b-4fa6-9331-bc2451586198','e6004ede-e6f3-4eba-afab-f7183802d64f','Next.js','uploads/1758875193_nextjs.png'),('1ca83bc8-a5de-4970-82d1-67c0e38ba3c6','d46303c5-91f0-4f7a-92a4-b6771a3209c2','CSS3','uploads/1758875274_css3.png'),('1e5b4eab-051e-4ec0-bff0-2779a3640d50','d46303c5-91f0-4f7a-92a4-b6771a3209c2','Bootstrap CSS','uploads/1758875274_bootstrap.png'),('2ca7cc63-6657-4f33-8479-b5d8c02d35c4','de6a79c3-2379-4e33-8b96-3566f7d7e593','Node.js','uploads/1758875320_nodejs.png'),('2d5c2c48-3e10-4176-94d2-7e535452d74f','de6a79c3-2379-4e33-8b96-3566f7d7e593','MongoDB','uploads/1758875320_mongodb.png'),('4a176566-664b-48b5-b85f-6029e0641a64','e6004ede-e6f3-4eba-afab-f7183802d64f','Zustand','uploads/1758875194_zustand.svg'),('76e7e8a0-4d49-4272-82ca-1b6969f31f9c','e6004ede-e6f3-4eba-afab-f7183802d64f','Vue.js','uploads/1758875193_vuejs.png'),('8076c929-4255-4480-82f4-05df4a3966f7','e6004ede-e6f3-4eba-afab-f7183802d64f','Tanstack Query','uploads/1758875193_tanstack.png'),('822f6e82-2caf-45d8-9c13-a30e9aea6d8c','2f0d08c0-9c77-41bf-bd1b-9426a6bf5454','JavaScript','uploads/1758873762_javascript.png'),('a3eed138-f3e8-41b9-a0a3-33f9ca5719a1','2f0d08c0-9c77-41bf-bd1b-9426a6bf5454','Golang','uploads/1758873762_golang.png'),('c8955ae7-8f2e-475f-92e0-af4845c62977','d46303c5-91f0-4f7a-92a4-b6771a3209c2','Ant Design','uploads/1758875274_Ant Design.png'),('c9569ee1-dbc3-4c03-afff-2b857ef0068c','de6a79c3-2379-4e33-8b96-3566f7d7e593','MySQL','uploads/1758875320_mysql.png'),('c99d2a52-f8a7-4db1-913e-2670a21a7b79','d46303c5-91f0-4f7a-92a4-b6771a3209c2','Shadcn UI','uploads/1758875274_shadcn-ui.png'),('d46a90dd-eda4-4d36-a331-f643367feac2','2f0d08c0-9c77-41bf-bd1b-9426a6bf5454','TypeScript','uploads/1758873762_typescript.png'),('d8f5d5c0-f4e9-4d4a-be31-c7b62a07baa4','2f0d08c0-9c77-41bf-bd1b-9426a6bf5454','HTML5','uploads/1758873762_html5.png'),('f59d5058-e2fc-491d-8629-cc312f462122','e6004ede-e6f3-4eba-afab-f7183802d64f','React.js','uploads/1758875193_react.png');
/*!40000 ALTER TABLE `skills` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'siroj05','$2a$10$E/qHnLmt5UpKZUM4wo0YZucODhhV6PjXzm8jQnlvoP9XqnTWXcITC');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'portfolio'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-29 10:07:03
