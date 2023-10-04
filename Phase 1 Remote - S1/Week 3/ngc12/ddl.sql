-- Create the authors table
CREATE TABLE authors (
    author_id INT PRIMARY KEY AUTO_INCREMENT,
    author_name VARCHAR(255) UNIQUE NOT NULL,
    author_country VARCHAR(255) NOT NULL
);

-- Create the books table
CREATE TABLE books (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) UNIQUE NOT NULL,
    author_id INT,
    publication_year INT,
    available_quantity INT,
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

-- Create the book_loans table
CREATE TABLE book_loans (
    loan_id INT PRIMARY KEY AUTO_INCREMENT,
    book_id INT,
    borrower_name VARCHAR(255) NOT NULL,
    loan_date DATE NOT NULL,
    return_date DATE,
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

-- Create an index on author_id in the books table
CREATE INDEX idx_author_id ON books (author_id);

-- Create an index on loan_date in the book_loans table
CREATE INDEX idx_loan_date ON book_loans (loan_date);   

-- Create an index on return_date in the book_loans table
CREATE INDEX idx_return_date ON book_loans (return_date);
