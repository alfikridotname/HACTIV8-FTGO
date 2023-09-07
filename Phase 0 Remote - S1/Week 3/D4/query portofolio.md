Create Database

    CREATE DATABASE portfolio;

Create Table biodata

    CREATE TABLE IF NOT EXISTS `biodata` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `nama` varchar(255) DEFAULT NOT NULL,
    `role` varchar(255) DEFAULT NULL,
    `availability` varchar(255) DEFAULT NULL,
    `age` int(11) DEFAULT NULL,
    `lokasi` varchar(255) DEFAULT NULL,
    `year_experience` int(11) DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL,
    `profile_photo` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
    );

Insert biodata

    INSERT INTO biodata ( nama, role, availability, age, lokasi, year_experience, email, profile_photo )
    VALUES
        (
            'Alfikri',
            'Backend Engineer', 
            'Fulltime Remote', 
            33, 
            'Padang, Indonesia', 
            10, 
            'alfikri.name@gmail.com', 
            'http://127.0.0.1:5500/Phase%200%20Remote%20-%20S1/Week%202/D2/portfolio/user.png' 
    );

Read biodata

    SELECT
        id,
        nama,
        role,
        availability,
        age,
        lokasi,
        year_experience,
        email,
        profile_photo 
    FROM
        biodata;


Update biodata

    UPDATE biodata SET availability = 'Partime Remote' WHERE id = 1;

Delete biodata

    DELETE FROM biodata WHERE id = 1;