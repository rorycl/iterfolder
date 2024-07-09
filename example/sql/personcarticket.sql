SELECT 
    *
FROM (
    SELECT 
        p.id AS personid
        ,p.firstname
        ,p.lastname
        ,c.registration
        ,c.manufacturer
        ,c.model
        ,left(t.uuid::text, 8) AS uuidpart
        ,date(t.datetime) as date
        ,count(c.registration) OvER (PARTITION BY c.registration) AS counter
    FROM
        people p
        JOIN cars c ON c.owner = p.id
        join tickets t ON t.car = c.registration
    ) x
ORDER BY
    counter desc
;
