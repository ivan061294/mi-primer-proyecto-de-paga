CREATE TABLE purchases (
  id           INTEGER     NOT NULL AUTO_INCREMENT COMMENT 'identificador de la compra',
  provider     INTEGER     NOT NULL COMMENT 'identificador del proveedor',
  product_id   VARCHAR(20) NOT NULL COMMENT 'identificador del producto',
  regdate      DATETIME    NOT NULL COMMENT 'fecha de la compra',
  CONSTRAINT pk_purchase PRIMARY KEY (id));