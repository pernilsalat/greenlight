ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime > 0);

ALTER TABLE movies ADD CONSTRAINT movies_year_check CHECK (year > 1888 AND year < date_part('year', now()));

ALTER TABLE movies ADD CONSTRAINT movies_genres_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);