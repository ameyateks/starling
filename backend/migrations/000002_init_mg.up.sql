CREATE TYPE direction AS ENUM ('IN','OUT');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    id UUID NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
    feed_item_uid VARCHAR(255) NOT NULL,
    category_uid VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    direction direction NOT NULL,
    transaction_time TIMESTAMPTZ NOT NULL,
    counter_party_name VARCHAR(255) NOT NULL, 
    counter_party_sub_entity_name VARCHAR(255),
    reference VARCHAR(255),
    spending_category VARCHAR(255) NOT NULL,
    user_note VARCHAR(255)
);