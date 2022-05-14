DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS mall;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS mall_customer;

CREATE TABLE user(
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL  
);

CREATE TABLE mall(
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL
);

CREATE TABLE item(
    id INTEGER PRIMARY KEY,
    mall_id INTEGER NOT NULL,
    total INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    data TEXT
);


CREATE TABLE mall_customer(
    mall_id INTEGER,
    user_id INTEGER,
    mall_item_id INTEGER,
    buy_count INTEGER
);