DROP TABLE IF EXISTS ppl;
CREATE TABLE ppl (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  email     VARCHAR(255) NOT NULL,
  phone      NUMERIC NOT NULL,
  msg   VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`)
);

