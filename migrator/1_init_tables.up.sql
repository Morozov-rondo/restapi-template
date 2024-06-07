/*
Создание файла миграции для таблиц базы данных. Пример использования:

CREATE TABLE IF NOT EXISTS search_history(
    id SERIAL PRIMARY KEY,
    query TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS address (
    id SERIAL PRIMARY KEY,
    location TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS history_search_address(
    id SERIAL PRIMARY KEY,
    search_id INTEGER REFERENCES search_history(id),
    address_id INTEGER REFERENCES address(id)
);

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;

 */



CREATE TABLE IF NOT EXISTS search_history(
    id SERIAL PRIMARY KEY,
    query TEXT UNIQUE NOT NULL
);
