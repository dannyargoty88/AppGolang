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