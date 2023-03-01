-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;'
    LANGUAGE plpgsql;

CREATE TABLE navigation
(
    id          UUID                 DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR     NOT NULL,
    code        VARCHAR,
    description TEXT                 DEFAULT NULL,
    sort        INT                  DEFAULT 500,
    parent_uuid UUID                 DEFAULT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON navigation
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO navigation
VALUES (uuid_generate_v4(), 'Пункт 1', '', '', 500),
       (uuid_generate_v4(), 'Пункт 2', 'item_2', '', 600),
       (uuid_generate_v4(), 'Пункт 3', 'item_3', '', 700),
       (uuid_generate_v4(), 'Пункт 4', 'item_4', '', 800),
       (uuid_generate_v4(), 'Пункт 5', 'item_5', '', 900),
       (uuid_generate_v4(), 'Пункт 6', 'item_6', '', 1000);

INSERT INTO navigation
VALUES (uuid_generate_v4(), 'Пункт 2.1', 'item_2_1', '', 500, (select id from navigation where code = 'item_2')),
       (uuid_generate_v4(), 'Пункт 2.2', 'item_2_2', '', 600, (select id from navigation where code = 'item_2')),
       (uuid_generate_v4(), 'Пункт 2.3', 'item_2_3', '', 700, (select id from navigation where code = 'item_2'));
-- +goose Down
DROP TABLE navigation;

DROP FUNCTION IF EXISTS trigger_set_timestamp();