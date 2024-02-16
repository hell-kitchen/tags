CREATE TABLE tags (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "color" VARCHAR(255) NOT NULL
);

---- create above / drop below ----

DROP TABLE tags;