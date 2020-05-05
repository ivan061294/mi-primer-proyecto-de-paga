CREATE TABLE invoicedetails (
  id          INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del detalle de la factura',
  invoice_id  INTEGER       NOT NULL COMMENT 'identificador de la factura',
  product_id  INTEGER       NOT NULL COMMENT 'identificador del producto',
  description VARCHAR(200)  NULL     COMMENT 'descripcion del producto',
  amount      INTEGER       NOT NULL COMMENT 'cantidad del producto',
  price       DECIMAL(10,2) NOT NULL COMMENT 'precio del producto',
  CONSTRAINT pk_invoicedetail PRIMARY KEY (id));