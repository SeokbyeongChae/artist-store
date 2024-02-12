CREATE TABLE `accounts` (
  `id` bigserial PRIMARY KEY,
  `salt` varchar(64) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchat(50) NOT NULL,
  `failed_login_count` tinyint,
  `last_login_at` timestamptz,
  `created_at` timestamptz DEFAULT (now())
);

CREATE TABLE `Session` (
  `id` bigserial PRIMARY KEY,
  `account_id` bigint,
  `refresh_token` varchar(255) NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `client_ip` varchar(255) NOT NULL,
  `is_blocked` bool NOT NULL,
  `expire_at` timestamptz NOT NULL,
  `created_at` timestamptz DEFAULT (now())
);

CREATE INDEX `accounts_index_0` ON `accounts` (`email`);

CREATE INDEX `accounts_index_1` ON `accounts` (`email`, `password`);

ALTER TABLE `Session` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
