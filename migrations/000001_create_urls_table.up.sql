CREATE TABLE IF NOT EXISTS `urls`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `short_url`   varchar(256) NOT NULL UNIQUE,
    `long_url`    varchar(256) NOT NULL UNIQUE,
    `description` varchar(256) NOT NULL,
    PRIMARY KEY (`id`)
);