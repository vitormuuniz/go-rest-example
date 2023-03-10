CREATE TABLE IF NOT EXISTS categories (
  id BIGINT PRIMARY KEY auto_increment,
  description VARCHAR(255) NOT NULL unique,
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP default CURRENT_TIMESTAMP
)
engine = InnoDB
default charset = utf8;

CREATE TABLE IF NOT EXISTS products (
  id BIGINT PRIMARY KEY auto_increment,
  name VARCHAR(512) NOT NULL unique,
  price DECIMAL(10,2) default 0.0,
  quantity INT(10) unsigned default 0,
  status CHAR(1) default 0,
  category_id BIGINT NOT NULL,
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP default CURRENT_TIMESTAMP,
  CONSTRAINT products_categories_fk FOREIGN KEY (category_id) REFERENCES categories(id)
  ON DELETE CASCADE ON UPDATE cascade
)
engine = InnoDB
default charset = utf8;