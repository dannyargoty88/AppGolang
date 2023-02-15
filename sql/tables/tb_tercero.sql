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