CREATE TABLE IF NOT EXISTS articles (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	markdown TEXT NOT NULL,
	slug TEXT NOT NULL UNIQUE,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
