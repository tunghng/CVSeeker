CREATE DATABASE IF NOT EXISTS TalentAcquisition;
USE TalentAcquisition;

CREATE TABLE `resumes` (
   `resume_id` int NOT NULL AUTO_INCREMENT,
   `full_text` text,
   `download_link` varchar(255) DEFAULT NULL,
   `vector_embedding` text,
   `created_at` datetime DEFAULT NULL,
   `updated_at` datetime DEFAULT NULL,
   PRIMARY KEY (`resume_id`)
);

CREATE TABLE `threads` (
   `id` varchar(100) NOT NULL,
   `created_at` datetime NOT NULL,
   `updated_at` datetime NOT NULL,
   `name` varchar(255) DEFAULT NULL,
   PRIMARY KEY (`id`)
);

CREATE TABLE `thread_resumes` (
  `thread_id` varchar(100) NOT NULL,
  `resume_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`thread_id`,`resume_id`),
  KEY `idx_thread_id` (`thread_id`),
  KEY `idx_resume_id` (`resume_id`)
);

CREATE TABLE `upload` (
  `id` int NOT NULL AUTO_INCREMENT,
  `document_id` varchar(255) DEFAULT NULL,
  `status` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
