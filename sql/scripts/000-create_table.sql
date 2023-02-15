--DROP TABLE IF EXISTS tb_usuario;

CREATE TABLE IF NOT EXISTS tb_usuario
(
    usuario_codigo serial NOT NULL,
	usuario_nombre character varying(150) NOT NULL,
	usuario_email character varying(50) NOT NULL,
	usuario_password character varying(200) NOT NULL,
    CONSTRAINT pk_tb_usuario PRIMARY KEY (usuario_codigo),
    UNIQUE (usuario_email)
);


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


-- DROP TABLE IF EXISTS public.tb_departamento;

CREATE TABLE IF NOT EXISTS public.tb_departamento
(
    departamento_codigo serial NOT NULL,
    departamento_nombre character varying(50) NOT NULL,
    CONSTRAINT pk_tb_departamento PRIMARY KEY (departamento_codigo)
);

-- DROP TABLE IF EXISTS public.tb_ciudad;

CREATE TABLE IF NOT EXISTS public.tb_ciudad
(
    ciudad_codigo serial NOT NULL,
    ciudad_nombre character varying(50) NOT NULL,
    departamento_codigo bigint NOT NULL,
    CONSTRAINT pk_tb_ciudad PRIMARY KEY (ciudad_codigo),
    CONSTRAINT fk_tb_ciudad_departamento_codigo FOREIGN KEY (departamento_codigo)
        REFERENCES public.tb_departamento (departamento_codigo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

--DROP TABLE IF EXISTS tb_tercero;

CREATE TABLE IF NOT EXISTS tb_tercero
(
    tercero_codigo serial NOT NULL,
    tercero_nombre character varying(50) NOT NULL,
	tercero_documento character varying(50) NOT NULL,
	ciudad_codigo bigint NOT NULL,
    CONSTRAINT pk_tb_tercero PRIMARY KEY (tercero_codigo),
	CONSTRAINT fk_tb_tercero_ciudad_codigo FOREIGN KEY (ciudad_codigo)
        REFERENCES public.tb_ciudad (ciudad_codigo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	UNIQUE (tercero_documento)
);


-- DROP TABLE IF EXISTS tb_ws;

CREATE TABLE IF NOT EXISTS tb_ws (
	ws_codigo bigint NOT NULL,
	ws_nombre character varying(200) NOT NULL,
	ws_url character varying(200) NOT NULL,
	ws_fechacreacion date NOT NULL,
	ws_fechamodificacion date NOT NULL,
	CONSTRAINT pk_tb_ws PRIMARY KEY (ws_codigo),
	UNIQUE (ws_nombre, ws_url)
);

-- DROP TABLE IF EXISTS tb_itemws;

CREATE TABLE IF NOT EXISTS tb_itemws (
	ws_codigo bigint NOT NULL,
	itews_codigo serial NOT NULL,	
	itews_method character varying(10) NOT NULL,
	itews_endpoint character varying(100) NOT NULL,
	itews_status character varying(10),
	itews_request text,
	itews_response text,
	itews_fecharequest timestamp,
	itews_fecharesponse timestamp,
	itews_fechacreacion date NOT NULL,
	itews_fechamodificacion date NOT NULL,
	CONSTRAINT pk_tb_itemws PRIMARY KEY (ws_codigo, itews_codigo),
	CONSTRAINT fk_tb_ws_ws_codigo FOREIGN KEY (ws_codigo)
        REFERENCES public.tb_ws (ws_codigo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);