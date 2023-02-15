-- DROP TABLE IF EXISTS public.tb_departamento;

CREATE TABLE IF NOT EXISTS public.tb_departamento
(
    departamento_codigo serial NOT NULL,
    departamento_nombre character varying(50) NOT NULL,
    CONSTRAINT pk_tb_departamento PRIMARY KEY (departamento_codigo)
);