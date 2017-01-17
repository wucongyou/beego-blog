CREATE TABLE blog (
  id      INT AUTO_INCREMENT,
  title   TEXT,
  content TEXT,
  created DATETIME,
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
