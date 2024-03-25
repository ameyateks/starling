CREATE TYPE direction AS ENUM ('IN','OUT');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    id UUID NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
    feedItemUid VARCHAR(255) NOT NULL,
    categoryUid VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    direction direction NOT NULL,
    transactionTime TIMESTAMPTZ NOT NULL,
    counterPartyName VARCHAR(255) NOT NULL, 
    counterPartySubEntityName VARCHAR(255),
    reference VARCHAR(255),
    spendingCategory VARCHAR(255) NOT NULL,
    userNote VARCHAR(255)
);