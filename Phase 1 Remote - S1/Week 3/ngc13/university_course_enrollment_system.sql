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

-- Insert sample students
INSERT INTO students (student_id, name, email, major, year_of_study)
VALUES
    (1, 'John Doe', 'johndoe@example.com', 'Computer Science', 2),
    (2, 'Jane Smith', 'janesmith@example.com', 'Mathematics', 3),
    (3, 'Alice Johnson', 'alice@example.com', 'Biology', 2),
    (4, 'Bob Brown', 'bob@example.com', 'Physics', 4),
    (5, 'Eva White', 'eva@example.com', 'History', 3);

-- Insert sample courses
INSERT INTO courses (course_id, title, instructor, schedule, credit_hours)
VALUES
    (101, 'Math', 'Prof. Smith', 'MWF 10:00 AM - 11:30 AM', 3),
    (102, 'Calculus I', 'Prof. Johnson', 'TTH 9:00 AM - 10:30 AM', 4),
    (103, 'Calculus II', 'Prof. Davis', 'MWF 1:00 PM - 2:30 PM', 3),
    (104, 'Calculus III', 'Prof. Brown', 'TTH 1:00 PM - 2:30 PM', 4),
    (105, 'Calculus IV', 'Prof. White', 'MWF 11:30 AM - 1:00 PM', 3);

-- Register students for courses with preferred schedules
INSERT INTO student_courses (student_id, course_id, preferred_schedule)
VALUES
    (1, 101, 'MWF 10:00 AM - 11:30 AM'),
    (1, 103, 'MWF 1:00 PM - 2:30 PM'),
    (2, 102, 'TTH 9:00 AM - 10:30 AM'),
    (3, 101, 'MWF 10:00 AM - 11:30 AM'),
    (4, 104, 'TTH 1:00 PM - 2:30 PM'),
    (5, 105, 'MWF 11:30 AM - 1:00 PM');

