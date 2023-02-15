
INSERT INTO tb_usuario VALUES 
(1, 'admin', 'admin@correo.com', '$2a$04$7OXFWP6olhP6MgtNEQd0b.kdHrnpSwwCzGSSCkHphhiz2htoTwXKS');-- Pass 123456
SELECT setval('tb_usuario_usuario_codigo_seq', (select max(usuario_codigo) from tb_usuario));

INSERT INTO tb_departamento VALUES 
(1, 'Bogotá DC'),
(2, 'Antioquia'),
(3, 'Valle del cauca');
SELECT setval('tb_departamento_departamento_codigo_seq', (select max(departamento_codigo) from tb_departamento));

INSERT INTO tb_ciudad VALUES 
(1,'Bogotá', 1),
(2,'Medellin', 2),
(3,'Cali', 3),
(4,'Yumbo', 3),
(5,'Palmira', 3);
SELECT setval('tb_ciudad_ciudad_codigo_seq', (select max(ciudad_codigo) from tb_ciudad));

INSERT INTO tb_tercero VALUES 
(1, 'Danny', '123456789', 3),
(2, 'Jhoan', '987654321', 1);
SELECT setval('tb_tercero_tercero_codigo_seq', (select max(tercero_codigo) from tb_tercero));


INSERT INTO tb_ws VALUES (1, 'Test Ws', 'https://pokeapi.co/api/v2/pokemon', current_date, current_date);
