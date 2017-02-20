CREATE DATABASE `test`
  DEFAULT CHARACTER SET utf8
  COLLATE utf8_general_ci;
USE `test`;
CREATE TABLE `blog` (
  `id`      INT                AUTO_INCREMENT
  COMMENT 'ID',
  `title`   TEXT COMMENT 'Title',
  `content` TEXT COMMENT 'Content',
  `ctime`   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT 'The create time',
  `mtime`   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  ON UPDATE CURRENT_TIMESTAMP
  COMMENT 'The latest update time',
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
