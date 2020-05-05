CREATE TABLE certifys (
	id           INTEGER       NOT NULL AUTO_INCREMENT COMMENT 'identificador de la acta de conformidad',
	quotation_id INTEGER       NOT NULL COMMENT 'identificador de la cotizacion',
	order_id     INTEGER       NOT NULL COMMENT 'identificador de la orden',
	regdate      DATETIME      NOT NULL COMMENT 'fecha de la acta de conformidad',
	description  VARCHAR(200)  NULL     COMMENT 'descripcion de la acta de conformidad',
	observation  VARCHAR(200)  NULL     COMMENT 'observacion de la acta de conformidad',
	CONSTRAINT pk_invoice PRIMARY KEY (id));