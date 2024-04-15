CREATE TYPE scrapeStatus AS ENUM ('SUCCESS', 'FAILURE');

CREATE TABLE scrape_logs AS (
    id UUID NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
    scrape_status scrapeStatus NOT NULL,
    time_of_run TIMESTAMPTZ NOT NULL
);