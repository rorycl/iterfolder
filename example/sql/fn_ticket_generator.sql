/* function to generate tickets */

CREATE OR REPLACE FUNCTION fn_tickets_generate (records int) RETURNS VOID AS $$

DECLARE
-- empty

    i INT = 1;
    latitude Decimal(8,6) = 51.513222;
    longitude Decimal(9,6) = -0.159015;

BEGIN

    LOOP
        EXIT WHEN i > records;

        INSERT INTO tickets VALUES (
            gen_random_uuid()
            ,(select registration from cars order by random() limit 1)
            ,(select date(d) FROM (
                    select timestamp '2023-01-01 00:00:00' + random() * 
                    (timestamp '2024-06-30 00:00:00' - timestamp '2023-01-01 00:00:00') AS d
              ) x)
            ,latitude
            ,longitude
        );

        RAISE NOTICE '%', i;
        i := i + 1;
        latitude := latitude + round(0.0004, 6);
        longitude := longitude - round(0.003, 6);
    END LOOP;

END;
$$ LANGUAGE plpgsql;
