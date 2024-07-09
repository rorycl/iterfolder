/*
postgres schema file for people/cars/tickets example
RCL 08 July 2024
*/

\set ON_ERROR_STOP on

DROP SCHEMA IF EXISTS iterf CASCADE;
CREATE SCHEMA iterf;
SET SEARCH_PATH = 'iterf';

CREATE TABLE people (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL
    ,firstname TEXT NOT NULL
    ,lastname TEXT NOT NULL
);

ALTER TABLE people ADD CONSTRAINT chk_names CHECK (
    length(firstname) >= 2
    AND 
    length(firstname) < 30
    AND
    length(lastname) >= 2
    AND 
    length(lastname) < 30
);

CREATE TABLE cars (
    registration TEXT NOT NULL PRIMARY KEY UNIQUE
    ,owner BIGINT NOT NULL REFERENCES people(id)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
    ,manufacturer TEXT NOT NULL
    ,model TEXT NOT NULL
);

CREATE TABLE tickets (
    uuid UUID NOT NULL PRIMARY KEY 
    ,car TEXT NOT NULL REFERENCES cars(registration)
    ,datetime TIMESTAMP WITH TIME ZONE NOT NULL
    ,lat Decimal(8,6) NOT NULL
    ,long Decimal(9,6) NOT NULL
);

