CREATE TABLE customers (
  id       INTEGER      NOT NULL AUTO_INCREMENT COMMENT 'identificador del cliente',
  fullname VARCHAR(120) NOT NULL COMMENT 'nombre del cliente',
  doctype  VARCHAR(3)   NOT NULL COMMENT 'tipo de documento (DNI, RUC)',
  docnum   VARCHAR(11)  NOT NULL COMMENT 'numero de docuemnto',
  address  VARCHAR(240) NOT NULL COMMENT 'direccion del cliente',
  CONSTRAINT pk_customer PRIMARY KEY (id)
);