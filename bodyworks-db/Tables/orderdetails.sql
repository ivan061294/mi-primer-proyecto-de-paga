CREATE TABLE orderdetails (
  id          INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del detalle de la orden de trabajo',
  order_id    INTEGER       NOT NULL COMMENT 'identificador de la orden de trabajo',
  description VARCHAR(200)  NULL     COMMENT 'descripcion del trabajo',
  workhours   INTEGER       NOT NULL COMMENT 'horas hombre',
  cloths      INTEGER       NOT NULL COMMENT 'Numero de paños',
  CONSTRAINT pk_orderdetail PRIMARY KEY (id));