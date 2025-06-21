ALTER USER cassandra WITH PASSWORD "test1234";

CREATE KEYSPACE IF NOT EXISTS football_team
WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE USER app WITH PASSWORD "test321" NOSUPERUSER;
GRANT ALL PERMISSIONS ON KEYSPACE football_team TO app;

USE football_team;

CREATE TABLE IF NOT EXISTS actual_team (
  nr INT PRIMARY KEY,
  name TEXT,
  last_name TEXT,
  position TEXT
);

CREATE TABLE IF NOT EXISTS incoming_transfers (
  nr INT PRIMARY KEY,
  name TEXT,
  last_name TEXT,
  position TEXT,
  price TEXT,
  team TEXT
);

CREATE TABLE IF NOT EXISTS outgoing_transfers (
  nr INT PRIMARY KEY,
  name TEXT,
  last_name TEXT,
  position TEXT,
  price TEXT,
  team TEXT
);