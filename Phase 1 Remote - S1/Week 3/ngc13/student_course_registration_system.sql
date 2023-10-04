-- Create the students table
CREATE TABLE students (
    student_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    major VARCHAR(255),
    year_of_study INT
);

-- Create the courses table
CREATE TABLE courses (
    course_id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    instructor VARCHAR(255),
    schedule VARCHAR(255),
    credit_hours INT
);

-- Create the student_courses table (Associative Entity)
CREATE TABLE student_courses (
    student_id INT,
    course_id INT,
    preferred_schedule VARCHAR(255),
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);
