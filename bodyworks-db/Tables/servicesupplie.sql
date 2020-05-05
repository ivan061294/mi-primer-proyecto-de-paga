CREATE TABLE servicesupplie (
  service_id INTEGER NOT NULL COMMENT 'identificador del servicio',
  supplie_id INTEGER NOT NULL COMMENT 'identificador del insumo',
  relation   INTEGER NOT NULL COMMENT 'relacion entre servicio insumo',
  rank       INTEGER NOT NULL COMMENT 'rango de la relacion'
);