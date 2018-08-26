CREATE TABLE user (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(190) NOT NULL COMMENT 'ニックネーム',
  age tinyint  NOT NULL COMMENT '年齢',
  status tinyint  NOT NULL DEFAULT 1,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

