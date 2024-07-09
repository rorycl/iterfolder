
\set ON_ERROR_STOP on

\i sql/schema.sql

\i sql/people.sql

\i sql/models_tmp.sql

\i sql/load_cars.sql

\i sql/fn_ticket_generator.sql

select * from fn_tickets_generate(1000);

