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