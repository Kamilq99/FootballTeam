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