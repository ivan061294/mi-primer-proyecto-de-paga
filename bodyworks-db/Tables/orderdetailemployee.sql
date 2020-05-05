CREATE TABLE orderdetailemployees (
  id             INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del detalle de la orden de trabajo',
  orderdetail_id INTEGER       NOT NULL COMMENT 'identificador de la orden de trabajo',
  employee_id    INTEGER       NULL     COMMENT 'identificador del empleado',
  workhours      INTEGER       NOT NULL COMMENT 'horas hombre',
  cloths         INTEGER       NOT NULL COMMENT 'Numero de paños',
  CONSTRAINT pk_orderdetailemployee PRIMARY KEY (id));