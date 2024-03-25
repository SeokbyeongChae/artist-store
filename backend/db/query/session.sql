/* name: CreateSession :execresult */
INSERT INTO sessions (
  uuid,
  account_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expire_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

/* name: GetSession :one */
SELECT * FROM sessions
WHERE uuid = ? LIMIT 1;
