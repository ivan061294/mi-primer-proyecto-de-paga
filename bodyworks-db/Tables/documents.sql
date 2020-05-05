CREATE TABLE documents (
	id           INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador del documento',
	name         VARCHAR(80)   NOT NULL COMMENT 'nombre del documento',
	content      BLOB          NULL     COMMENT 'contenido del documento',
	CONSTRAINT pk_document PRIMARY KEY (id));