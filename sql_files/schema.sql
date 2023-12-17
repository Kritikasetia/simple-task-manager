-- Creating the task manager database
CREATE DATABASE IF NOT EXISTS taskmanagerdb;

-- Use the database
USE taskmanagerdb;

-- Creating  the tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT
);
