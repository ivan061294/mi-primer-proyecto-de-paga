CREATE TABLE settings (
  id          INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador de la configuracion',
  name        VARCHAR(80)   NOT NULL COMMENT 'llave de la configuracion',
  value       VARCHAR(180)  NOT NULL COMMENT 'valor de la configuracion',
  regdate     DATETIME      NOT NULL COMMENT 'fecha de registro',
  CONSTRAINT pk_setting PRIMARY KEY (id));