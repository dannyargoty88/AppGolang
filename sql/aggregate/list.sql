-- DROP FUNCTION IF EXISTS public.fc_convertir_a_lista(text, text);

CREATE OR REPLACE FUNCTION public.fc_convertir_a_lista(
	text,
	text)
    RETURNS text
    LANGUAGE 'sql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
select case
    WHEN $2 is null or $2 = '' THEN $1
    WHEN $1 is null or $1 = '' THEN $2
   ELSE $1 || ', ' || $2
  END
$BODY$;
	
			
-- Aggregate: list;

-- DROP AGGREGATE IF EXISTS public.list(text);

CREATE OR REPLACE AGGREGATE public.list(text) (
    SFUNC = fc_convertir_a_lista,
    STYPE = text ,
    FINALFUNC_MODIFY = READ_ONLY,
    MFINALFUNC_MODIFY = READ_ONLY
);