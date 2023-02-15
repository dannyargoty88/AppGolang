-- DROP TABLE IF EXISTS tb_tipows;

CREATE TABLE IF NOT EXISTS tb_tipows (
	tipws_codigo bigint NOT NULL,
	tipws_nombre character varying(50) NOT NULL,
	tipws_fechacreacion date NOT NULL,
	tipws_fechamodificacion date NOT NULL,
	CONSTRAINT pk_tb_tipows PRIMARY KEY (tipws_codigo)
);