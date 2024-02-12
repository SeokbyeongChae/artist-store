/* name: GetAccount :one */
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

/* name: ListAccounts :many */
SELECT * FROM accounts
ORDER BY id
LIMIT ?
OFFSET ?;

/* name: CreateAccount :execresult */
INSERT INTO accounts (
  salt,
  email,
  password
) VALUES (
  ?, ?, ?
);

/* name: UpdateAccount :exec */
UPDATE accounts
SET balance = ?
WHERE id = ?;

/* name: DeleteAccount :exec */
DELETE FROM accounts 
WHERE id = ?;