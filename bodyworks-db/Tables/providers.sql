CREATE TABLE providers (
  id           INTEGER      NOT NULL AUTO_INCREMENT COMMENT 'identificador del proveedor',
  name         VARCHAR(120) NOT NULL COMMENT 'nombre del proveedor',
  status       VARCHAR(10)  NOT NULL COMMENT 'Estado del proveedor (Activo, Desactivo)',
  regdate      DATETIME     NOT NULL COMMENT 'fecha de la compra',
  CONSTRAINT pk_purchase PRIMARY KEY (id));