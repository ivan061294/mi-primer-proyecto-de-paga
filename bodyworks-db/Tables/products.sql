CREATE TABLE products (
  id          INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del producto',
  description VARCHAR(200)  NOT NULL COMMENT 'descripcion del producto',
  measurement VARCHAR(80)   NOT NULL COMMENT 'unidad de medida del producto',
  unitprice   DECIMAL(10,2) NOT NULL COMMENT 'precio unitario del producto',
  type        VARCHAR(40)       NULL COMMENT 'tipo del producto (INSUMO, SERVICIO)',
  category    VARCHAR(40)       NULL COMMENT 'categoria del producto (PINTURA)',
  stock       INTEGER       NOT NULL COMMENT 'numero de items disponibles',
  CONSTRAINT pk_product PRIMARY KEY (id));