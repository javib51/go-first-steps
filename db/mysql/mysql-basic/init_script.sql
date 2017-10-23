DROP DATABASE IF EXISTS hello;
CREATE DATABASE hello;
USE hello;

/* Tables definition */
CREATE TABLE messages ( 
	id INT NOT NULL AUTO_INCREMENT,
	message VARCHAR(100) NOT NULL,
	PRIMARY KEY(id));

/* Users Definition */
GRANT ALL PRIVILEGES ON hello.* TO 'gouser'@'%' IDENTIFIED BY "gupw55"; 
FLUSH PRIVILEGES;
