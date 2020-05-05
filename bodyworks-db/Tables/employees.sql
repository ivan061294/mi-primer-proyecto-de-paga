CREATE TABLE employees (
  id       INTEGER     NOT NULL AUTO_INCREMENT COMMENT 'identificador del empleado',
  name     VARCHAR(80) NOT NULL COMMENT 'nombre del empleado',
  lastname VARCHAR(80) NOT NULL COMMENT 'apellido del empleado',
  care     VARCHAR(80) NOT NULL COMMENT 'cargo del empĺeado',
  CONSTRAINT pk_employee PRIMARY KEY (id));