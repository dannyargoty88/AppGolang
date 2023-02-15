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