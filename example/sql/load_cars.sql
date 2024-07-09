/*
tmp
*/

\set NUMROWS 1000
-- 250 models : modulo starts at 0
\set NUMMODELS 249 
-- 102 people : modulo starts at 0
\set NUMPEOPLE 101 

INSERT INTO cars
SELECT
    *
FROM (
    SELECT
        x.char_prefix || '-' || x.numstring || '-' || left(x.char_prefix, 1) AS reg
        ,p.id as person
        ,m.manufacturer
        ,m.model
        /*
        ,x.counter AS counter
        ,m.rn AS mcounter
        ,(x.counter % :NUMMODELS) AS mod
        ,(x.counter % :NUMPEOPLE) AS mod2
        */
    FROM (
        SELECT
            i AS counter
            ,lpad(concat('', i), 5, '0') AS numstring
            ,left(upper(regexp_replace(md5(random()::text), '[0-9]', '', 'g')), 3) AS char_prefix
        FROM
            generate_series(1, :NUMROWS) i
    ) x
    LEFT JOIN (
        SELECT
            row_number() OVER () AS rn
            ,*
        FROM
            manu_model
        ORDER BY
            RANDOM()
     ) m ON ((x.counter % :NUMMODELS)+1 = m.rn)
    LEFT JOIN (
        SELECT
            row_number() OVER () AS rn
            ,id
        FROM
            people
        ORDER BY
            random()
    ) p ON ((x.counter % :NUMPEOPLE)+1 = p.rn)
) y
ORDER BY person
;
