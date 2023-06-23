UPDATE users
SET email = CONCAT_WS('_', email, id)
WHERE email IN (
   SELECT email
   FROM users
   GROUP BY email
   HAVING COUNT(*) > 1
);
