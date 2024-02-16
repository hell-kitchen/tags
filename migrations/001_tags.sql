CREATE TABLE tags
(
    "id"    UUID PRIMARY KEY NOT NULL UNIQUE,
    "name"  VARCHAR(255)     NOT NULL UNIQUE,
    "slug"  VARCHAR(255)     NOT NULL UNIQUE,
    "color" VARCHAR(255)     NOT NULL UNIQUE
);

---- create above / drop below ----

DROP TABLE tags;