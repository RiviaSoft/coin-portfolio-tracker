-- MYSQL
DROP DATABASE IF EXISTS portfolio_tracker_db;
CREATE DATABASE portfolio_tracker_db;
USE portfolio_tracker_db;

CREATE TABLE `recent_buy_operations` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL,
  `coin_symbol` varchar(10) NOT NULL,
  `coin_amount` double NOT NULL,
  `buy_cost` double NOT NULL,
  PRIMARY KEY (id)
) CHARSET=utf8mb4;

CREATE TABLE `archived_operations` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL,
  `coin_symbol` varchar(10) NOT NULL,
  `coin_amount` double NOT NULL,
  `buy_cost` double NOT NULL,
  `sell_cost` double NOT NULL,
    PRIMARY KEY (id)
) CHARSET=utf8mb4;

CREATE TABLE `users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL,
  `email` varchar(40) NOT NULL UNIQUE,
  `password_hash` varchar(300) NOT NULL,
    PRIMARY KEY (id)
)  CHARSET=utf8mb4;
