DROP DATABASE IF EXISTS AdminLogin;
CREATE DATABASE AdminLogin;
USE AdminLogin;
CREATE TABLE users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username varchar(64) NOT NULL,
    password varchar(128) NOT NULL
);

