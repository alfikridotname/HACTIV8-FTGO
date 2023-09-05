Create Database

    CREATE DATABASE `tourism`;

Create Table Destinations

    CREATE TABLE `destinations` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `destination_name` varchar(255) NOT NULL,
      `description` text,
      `location` varchar(255) DEFAULT NULL,
      `rating` decimal(3,2) DEFAULT NULL,
      PRIMARY KEY (`id`)
    );

Create Table Hotels

    CREATE TABLE `hotels` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `hotel_name` varchar(255) NOT NULL,
      `destination_id` int(11) DEFAULT NULL,
      `description` text,
      `address` varchar(255) DEFAULT NULL,
      `rating` decimal(3,2) DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `destination_id` (`destination_id`),
      CONSTRAINT `hotels_ibfk_1` FOREIGN KEY (`destination_id`) REFERENCES `destinations` (`id`)
    );

Create Table Guests

    CREATE TABLE `guests` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `first_name` varchar(255) DEFAULT NULL,
        `last_name` varchar(255) DEFAULT NULL,
        PRIMARY KEY (`id`)
    );

Create Table Bookings

    CREATE TABLE `bookings` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `guest_id` int(11) NOT NULL,
      `check_in_date` date DEFAULT NULL,
      `check_out_date` date DEFAULT NULL,
      `hotel_id` int(11) DEFAULT NULL,
      `total_price` decimal(10,2) DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `hotel_id` (`hotel_id`),
      CONSTRAINT `bookings_ibfk_1` FOREIGN KEY (`hotel_id`) REFERENCES `hotels` (`id`)
    );

Destination

    SELECT * FROM destinations;

Hotels
    
    SELECT * FROM hotels;

Bookings
    
    SELECT * FROM bookings;

Hotel By Destination ID
    
    SELECT * FROM hotels WHERE destination_id = 1;
    
Average Rating By Destination ID  
    
    SELECT ROUND(AVG(rating),2) AS average_rating
    
    FROM Hotels
    
    WHERE destination_id = 1;

Tampil Booking Details

    SELECT
        b.id,
        CONCAT(g.first_name,g.last_name) AS guest,
        h.hotel_name,
        h.rating,
        d.destination_name
    FROM
        bookings b
        INNER JOIN guests g ON b.guest_id = g.id
        INNER JOIN hotels h ON b.hotel_id = h.id
        INNER JOIN destinations d ON h.destination_id = d.id