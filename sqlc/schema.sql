CREATE TABLE upstreams (
  id SERIAL PRIMARY KEY,
  name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR UNIQUE NOT NULL,
  username VARCHAR UNIQUE NOT NULL,
  email VARCHAR UNIQUE NOT NULL
);

CREATE TABLE users_upstreams (
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  upstream_id INTEGER NOT NULL REFERENCES upstreams(id) ON DELETE CASCADE
);

CREATE TABLE targets (
  id SERIAL PRIMARY KEY,
  protocol VARCHAR NOT NULL,
  host VARCHAR NOT NULL,
  port INTEGER NOT NULL,
  upstream_id INTEGER NOT NULL REFERENCES upstreams(id) ON DELETE CASCADE
);

CREATE TABLE routes (
  id SERIAL PRIMARY KEY,
  path VARCHAR UNIQUE NOT NULL,
  private BOOLEAN NOT NULL,
  inner_path VARCHAR,
  upstream_id INTEGER NOT NULL REFERENCES upstreams(id) ON DELETE CASCADE
);

CREATE TABLE api_consumers (
  id SERIAL PRIMARY KEY,
  name VARCHAR UNIQUE NOT NULL,
  api_key VARCHAR UNIQUE NOT NULL
);

CREATE TABLE api_consumers_routes (
  api_consumer_id INTEGER NOT NULL REFERENCES api_consumers(id) ON DELETE CASCADE,
  route_id INTEGER NOT NULL REFERENCES routes(id) ON DELETE CASCADE,
  PRIMARY KEY(api_consumer_id, route_id)
)
