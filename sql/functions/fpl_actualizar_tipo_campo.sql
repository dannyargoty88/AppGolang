-- FUNCTION: public.fpl_actualizar_tipo_campo(character varying, character varying)

-- DROP FUNCTION IF EXISTS public.fpl_actualizar_tipo_campo(character varying, character varying);

CREATE OR REPLACE FUNCTION public.fpl_actualizar_tipo_campo(
	character varying,
	character varying)
    RETURNS character varying
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$


DECLARE 

	p_campo  	ALIAS FOR $1;
	p_tipo_cambio 	ALIAS FOR $2;
	
	v_record RECORD; 
	v_retvalor character varying;

	sql_borrar_vista  	character varying;
	sql_crear_vista 	character varying;
	sql_actualizar_tabla  	character varying;
  
BEGIN

	v_retvalor := 'Ok';
	sql_crear_vista:= '';

	/* **************** CONSULTA Y BORRA LAS VISTAS **************/

	sql_borrar_vista := 'SELECT ''DROP VIEW ''||viewname||'';'' as borrar_vista,
				''CREATE OR REPLACE VIEW ''||viewname||'' AS 
						''||definition as crear_vista,  
			       viewname as vista
			       FROM pg_views 
			       WHERE definition ilike ''%'||p_campo||'%''';

	RAISE NOTICE 'sql_borrar_vista: %',sql_borrar_vista;
	
	FOR v_record IN EXECUTE sql_borrar_vista LOOP
		RAISE NOTICE 'Vista Borrada: %',v_record.borrar_vista;
		sql_crear_vista := sql_crear_vista||' '||v_record.crear_vista;
		EXECUTE v_record.borrar_vista;
	END LOOP;

	/* **************** CONSULTA Y ACTUALIZA LAS TABLAS **************/

	sql_actualizar_tabla := 'SELECT ''ALTER TABLE ''||table_name||'' ALTER COLUMN ''||column_name||'' TYPE '||p_tipo_cambio||';'' as actualiza_tabla
	               FROM information_schema.columns 
	               WHERE column_name ilike '''||p_campo||'%''';

	RAISE NOTICE 'sql_actualizar_tabla: %',sql_actualizar_tabla;
	
	FOR v_record IN EXECUTE sql_actualizar_tabla LOOP
		RAISE NOTICE 'Tabla Actualizada: %',v_record.actualiza_tabla;
		EXECUTE v_record.actualiza_tabla;
	END LOOP;


	/* **************** CREA LAS VISTAS **************/

	RAISE NOTICE 'sql_crear_vista: %',sql_crear_vista;
	EXECUTE sql_crear_vista;

	RETURN v_retvalor;
  
END; 
$BODY$;