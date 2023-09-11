-- Query 1: List the titles of books borrowed by a member (identified by a specific Member_ID).
SELECT b.title
FROM lending_transaction lt
JOIN book b ON lt.isbn = b.isbn
WHERE lt.member_id = 1;

-- Query 2: Find out the members who have borrowed books but have not returned yet.
SELECT DISTINCT m.member_id, m.frist_name, m.last_name
FROM member m
JOIN lending_transaction lt ON m.member_id = lt.member_id
WHERE lt.date_of_return IS NULL;

-- Query 3: Find out the total number of books borrowed by a member (identified by a specific Member_ID).
SELECT m.member_id, m.frist_name, m.last_name, COUNT(lt.transaction_id) AS total_borrowed
FROM member m
LEFT JOIN lending_transaction lt ON m.member_id = lt.member_id
WHERE m.member_id = 1
GROUP BY m.member_id;

-- Query 4: List the books which were not returned in good condition.
SELECT b.title, lt.condition_at_return
FROM lending_transaction lt
JOIN book b ON lt.isbn = b.isbn
WHERE lt.condition_at_return <> 'Good';

-- Query 5: Identify members who have borrowed more than one book at a time.
SELECT m.member_id, m.frist_name, m.last_name
FROM member m
JOIN (
    SELECT member_id
    FROM lending_transaction
    GROUP BY member_id
    HAVING COUNT(*) > 1
) multiple_borrowers ON m.member_id = multiple_borrowers.member_id;