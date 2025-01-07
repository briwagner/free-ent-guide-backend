CREATE TABLE IF NOT EXISTS caches (
  id BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(256),
  value MEDIUMBLOB,
  updated_at datetime,
  expires datetime,
  UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS mlb_teams (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  updated_at datetime,
  team_id BIGINT NOT NULL,
  name VARCHAR(256),
  link VARCHAR(256),
  UNIQUE(team_id)
);

CREATE TABLE IF NOT EXISTS mlb_games (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  updated_at datetime,
  gametime datetime,
  game_id BIGINT NOT NULL,
  description text,
  status text,
  link text,
  home_id BIGINT NOT NULL,
  visitor_id BIGINT NOT NULL,
  home_score BIGINT,
  visitor_score BIGINT,
  UNIQUE `idx_gameid_gametime` (`game_id`, `gametime`),
  FOREIGN KEY (home_id) REFERENCES mlb_teams(id),
  FOREIGN KEY (visitor_id) REFERENCES mlb_teams(id)
);

CREATE TABLE IF NOT EXISTS nhl_teams (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  updated_at datetime,
  team_id BIGINT NOT NULL,
  franchise_id BIGINT NOT NULL,
  tricode VARCHAR(64),
  name VARCHAR(256),
  link VARCHAR(256),
  UNIQUE(team_id)
);

CREATE TABLE IF NOT EXISTS nhl_games (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  updated_at datetime,
  gametime datetime,
  game_id BIGINT,
  description text,
  status text,
  link text,
  home_id BIGINT NOT NULL,
  visitor_id BIGINT NOT NULL,
  home_score BIGINT,
  visitor_score BIGINT,
  UNIQUE `idx_gameid_gametime` (`game_id`, `gametime`),
  FOREIGN KEY (home_id) REFERENCES nhl_teams(id),
  FOREIGN KEY (visitor_id) REFERENCES nhl_teams(id)
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(256) NOT NULL,
  password_hash text NOT NULL,
  reset_code text,
  data JSON,
  created_at datetime,
  updated_at datetime,
  status BOOL DEFAULT 1,
  UNIQUE(email)
);