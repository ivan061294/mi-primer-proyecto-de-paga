CREATE TABLE orders (
  id           INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador de la orden de trabajo',
  quotation_id INTEGER       NOT NULL COMMENT 'identificador de la cotizacion',
  worktype     VARCHAR(15)   NOT NULL COMMENT 'tipo de trabajo de la orden de trabajo (PDI, SINIESTRO, UNIDAD COMPLETA, PIEZAS)',
  status       VARCHAR(20)   NOT NULL COMMENT 'estado de la orden de trabajo (E: EN PROCESO, T: TERMINADO)',
  totalhours   INTEGER       NOT NULL COMMENT 'total de horas hombre',
  totalcloths  INTEGER       NOT NULL COMMENT 'total de paños',
  startdate    DATETIME      NOT NULL COMMENT 'fecha de inicio de la orden de trabajo',
  enddate      DATETIME      NOT NULL COMMENT 'fecha de entrega de la orden de trabajo',
  regdate      DATETIME      NOT NULL COMMENT 'fecha de registro de la orden de trabajo',
  CONSTRAINT pk_order PRIMARY KEY (id));