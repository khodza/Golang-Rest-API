UPDATE users
SET email = CONCAT(email,"_",id)
WHERE email IN(
    SELECT email
    FROM users
    GROUP BY email
    HAVING COUNT (*) > 1
)