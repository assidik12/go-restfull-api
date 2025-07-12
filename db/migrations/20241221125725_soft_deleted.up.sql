CREATE TABLE `soft_deleted` (
    `id_transaction` int(11) NOT NULL,
    `deleted_at` datetime DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;