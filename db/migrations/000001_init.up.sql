CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    irina_id VARCHAR(50) UNIQUE NOT NULL,
    nickname VARCHAR(50) UNIQUE NOT NULL,
    wins_count INTEGER NOT NULL DEFAULT 0,
    losses_count INTEGER NOT NULL DEFAULT 0,
    leaves_count INTEGER NOT NULL DEFAULT 0,
    rating INTEGER NOT NULL,
    rank_position INTEGER NOT NULL,
    winrate DECIMAL(5,2) NOT NULL,
    total_games INTEGER NOT NULL DEFAULT 0,
    broken_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_players_irina_id ON players(irina_id);
CREATE INDEX IF NOT EXISTS idx_players_nickname ON players(nickname);
CREATE INDEX IF NOT EXISTS idx_players_rating ON players(rating DESC);