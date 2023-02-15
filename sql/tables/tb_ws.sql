-- DROP TABLE IF EXISTS tb_ws;

CREATE TABLE IF NOT EXISTS tb_ws (
	ws_codigo bigint NOT NULL,
	tipws_codigo bigint NOT NULL,
	ws_nombre character varying(200) NOT NULL,
	ws_url character varying(200) NOT NULL,
	ws_fechacreacion date NOT NULL,
	ws_fechamodificacion date NOT NULL,
	CONSTRAINT pk_tb_ws PRIMARY KEY (ws_codigo),
	UNIQUE (ws_nombre, ws_url)
);