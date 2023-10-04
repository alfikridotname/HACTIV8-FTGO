-- Create the students table
CREATE TABLE students (
    student_id INT PRIMARY KEY,
    student_name VARCHAR(100) NOT NULL
);

-- Create the professors table
CREATE TABLE professors (
    professor_id INT PRIMARY KEY,
    professor_name VARCHAR(100) NOT NULL
);

-- Create the courses table
CREATE TABLE courses (
    course_id INT PRIMARY KEY,
    course_title VARCHAR(200) NOT NULL,
    max_capacity INT
);

-- Create the Enrollment ternary relationship table
CREATE TABLE Enrollment (
    enrollment_id INT PRIMARY KEY,
    student_id INT,
    course_id INT,
    enrollment_date DATE,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

-- Create the TeachingAssignment ternary relationship table
CREATE TABLE TeachingAssignment (
    assignment_id INT PRIMARY KEY,
    professor_id INT,
    course_id INT,
    start_date DATE,
    FOREIGN KEY (professor_id) REFERENCES professors(professor_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

-- Insert sample students
INSERT INTO students (student_id, student_name)
VALUES
    (1, 'John Doe'),
    (2, 'Jane Smith'),
    (3, 'Alice Johnson');

-- Insert sample professors
INSERT INTO professors (professor_id, professor_name)
VALUES
    (101, 'Prof. Anderson'),
    (102, 'Prof. Brown'),
    (103, 'Prof. Clark'),
    (104, 'Prof. Davis'),
    (105, 'Prof. Evans');

-- Insert sample courses
INSERT INTO courses (course_id, course_title, max_capacity)
VALUES
    (201, 'Introduction to Computer Science', 30),
    (202, 'Calculus I', 40),
    (203, 'Biology 101', 35);

SELECT students.student_name
FROM students
JOIN Enrollment ON students.student_id = Enrollment.student_id
WHERE Enrollment.course_id = 201;

SELECT courses.course_title
FROM courses
JOIN TeachingAssignment ON courses.course_id = TeachingAssignment.course_id
WHERE TeachingAssignment.professor_id = 101;

SELECT professors.professor_name
FROM professors
JOIN TeachingAssignment ON professors.professor_id = TeachingAssignment.professor_id
WHERE TeachingAssignment.course_id = 202;
