CREATE TABLE quotedetails (
  id           INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del detalle de la cotizacion',
  quotation_id INTEGER       NOT NULL COMMENT 'identificador de la cotizacion',
  product_id   INTEGER       NOT NULL COMMENT 'identificador del producto',
  description  VARCHAR(200)  NULL     COMMENT 'Descipcion del producto',
  amount       INTEGER       NOT NULL COMMENT 'cantidad del producto',
  price        DECIMAL(10,2) NOT NULL COMMENT 'precio del producto',
  CONSTRAINT pk_quotedetail PRIMARY KEY (id));