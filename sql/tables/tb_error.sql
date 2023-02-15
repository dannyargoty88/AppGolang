-- DROP TABLE IF EXISTS tb_error;

CREATE TABLE IF NOT EXISTS tb_error (
	error_codigo serial NOT NULL,
	error_fecha timestamp NOT NULL,
	error_ip character varying(200) NOT NULL,
	error_host character varying(200) NOT NULL,
	error_useragent character varying(200) NOT NULL,
	error_method character varying(10) NOT NULL,
	error_url character varying(200) NOT NULL,
	error_status character varying(4) NOT NULL,
	error_message text NOT NULL,
	CONSTRAINT pk_tb_error PRIMARY KEY (error_codigo)
);