[Users]

CREATE TABLE Users (
    ->     user_id INT AUTO_INCREMENT PRIMARY KEY,
    ->     username VARCHAR(50) NOT NULL,
    ->     password VARCHAR(255) NOT NULL,
    ->     role ENUM('admin', 'user') NOT NULL,
    ->     first_name VARCHAR(50) NOT NULL,
    ->     last_name VARCHAR(50) NOT NULL,
    ->     email VARCHAR(100) NOT NULL UNIQUE,
    ->     phone_number VARCHAR(15),
    ->     date_of_birth DATE,
    ->     address VARCHAR(255)
    -> );


[Doctors]

CREATE TABLE Doctors (
    ->     doctor_id INT AUTO_INCREMENT PRIMARY KEY,
    ->     name VARCHAR(100) NOT NULL,
    ->     specialty VARCHAR(100) NOT NULL,
    ->     availability_timings VARCHAR(100) NOT NULL 
    -> );


[Appointments]

CREATE TABLE Appointments (
    ->     appointment_id INT AUTO_INCREMENT PRIMARY KEY,
    ->     user_id INT NOT NULL,
    ->     doctor_id INT NOT NULL,
    ->     appointment_date DATETIME NOT NULL,
    ->     FOREIGN KEY (user_id) REFERENCES Users(user_id),
    ->     FOREIGN KEY (doctor_id) REFERENCES Doctors(doctor_id)
    -> );

