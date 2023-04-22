DROP TABLE articles;

CREATE Table articles (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	markdown TEXT NOT NULL,
	slug TEXT NOT NULL UNIQUE,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO articles (title, description, markdown, slug)
VALUES 
	('My first article', 'My first description', '# Welcome to article 1', 'my-first-article'),
	('My second article', 'My second description', '# Welcome to article 2', 'my-second-article'),
	('My third article', 'My third description', '# Welcome to article 3', 'my-third-article'),
	('My fourth article', 'My fourth description', '# Welcome to article 4', 'my-fourth-article'),
	('My fifth article', 'My fifth description', '# Welcome to article 5', 'my-fifth-article'),
	('My sixth article', 'My sixth description', '# Welcome to article 6', 'my-sixth-article');