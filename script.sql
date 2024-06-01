CREATE TABLE urls (
  id         SERIAL PRIMARY KEY,
  original_url     VARCHAR(128) NOT NULL,
  shorter_url     VARCHAR(255) NOT NULL,
  created_at      TIMESTAMP NOT NULL
);