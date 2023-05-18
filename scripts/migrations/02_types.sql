CREATE TABLE `types`
(
    `id`         INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    `title`      VARCHAR(255)                    NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at`  TIMESTAMP,
    PRIMARY KEY (`id`)
);