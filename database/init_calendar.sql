CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Tasks
(
    id          int GENERATED ALWAYS AS IDENTITY,
    display_id  uuid                  DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
    name        varchar(64) NOT NULL UNIQUE,
    description varchar(126) NOT NULL UNIQUE,
    starts_at   timestamptz  DEFAULT current_timestamp NOT NULL,
    finishes_at timestamptz  DEFAULT current_timestamp NOT NULL,
    priority    smallint     NOT NULL,
    color       smallint     NOT NULL,
    completed   boolean      NOT NULL default FALSE,
    PRIMARY KEY (id)
);
CREATE INDEX idx_tasks ON Tasks (id DESC, display_id);