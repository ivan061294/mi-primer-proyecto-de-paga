CREATE TABLE users (
  id         INTEGER      NOT NULL AUTO_INCREMENT COMMENT 'identificador del usuario',
  username   VARCHAR(16)  NOT NULL COMMENT 'nombre del usuario',
  email      VARCHAR(255) NULL COMMENT 'correo del usuario',
  password   VARCHAR(32)  NOT NULL COMMENT 'contraseña del usuario',
  createtime TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'fecha de creacion',
  CONSTRAINT pk_user PRIMARY KEY (id));