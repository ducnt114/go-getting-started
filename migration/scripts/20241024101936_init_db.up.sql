DROP TABLE IF EXISTS profile;
DROP TABLE IF EXISTS book;
DROP TABLE IF EXISTS user;

CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `age` bigint DEFAULT NULL,
  `pass` longtext,
  `tags` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `book` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `title` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_books` (`user_id`),
  CONSTRAINT `fk_user_books` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `profile` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `user_name` varchar(50) DEFAULT NULL,
  `bio` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (1, '2024-10-07 12:03:06.501', '2024-10-07 12:03:06.501', null, 'duc', 32, null, null);
INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (2, '2024-10-14 12:24:38.984', '2024-10-14 12:24:38.984', null, 'nam', 22, null, null);
INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (3, '2024-10-14 12:31:24.143', '2024-10-14 12:31:24.143', null, 'long', 33, null, null);
INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (4, '2024-10-14 12:39:12.016', '2024-10-14 12:39:12.016', null, 'anh', 44, null, null);
INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (5, '2024-10-14 12:39:21.441', '2024-10-14 12:39:21.441', null, 'nguyen', 55, null, null);
INSERT INTO user (id, created_at, updated_at, deleted_at, name, age, pass, tags) VALUES (6, '2024-10-24 03:18:26.998', '2024-10-24 03:18:26.998', null, 'quang', 42, 'random-pass-123', '[{"key":"key-1","value":"val-1"},{"key":"key-2","value":"val-2"}]');

INSERT INTO book (id, name, title, user_id) VALUES (1, 'ten sach 1', 'tieu de sach 1', 1);
INSERT INTO book (id, name, title, user_id) VALUES (2, 'ten sach 2', 'tieu de sach 2', 1);
INSERT INTO book (id, name, title, user_id) VALUES (3, 'ten sach 3', 'tieu de sach 3', 3);
INSERT INTO book (id, name, title, user_id) VALUES (4, 'ten sach 4', 'tieu de sach 4', 2);
INSERT INTO book (id, name, title, user_id) VALUES (5, 'ten sach 5', 'tieu de sach 5', 1);

INSERT INTO profile (id, created_at, updated_at, deleted_at, user_name, bio) VALUES (1, '2024-10-24 10:13:37', '2024-10-24 10:13:38', null, 'nam', 'bio nam');
INSERT INTO profile (id, created_at, updated_at, deleted_at, user_name, bio) VALUES (2, '2024-10-24 10:14:08', '2024-10-24 10:14:10', null, 'long', 'bio long');
INSERT INTO profile (id, created_at, updated_at, deleted_at, user_name, bio) VALUES (3, '2024-10-24 10:14:19', '2024-10-24 10:14:20', null, 'duc', 'bio duc');
