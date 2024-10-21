-- create_climate_schema.up.sql
CREATE TABLE places
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(100)  NOT NULL,
    country   VARCHAR(100)  NOT NULL,
    latitude  DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL,
    altitude  INTEGER
);

CREATE TABLE climate_normals
(
    id           SERIAL PRIMARY KEY,
    place_id     INTEGER REFERENCES places (id),
    month        INTEGER NOT NULL CHECK (month BETWEEN 1 AND 12),
    tavg         DECIMAL(5, 2),
    tmin         DECIMAL(5, 2),
    tmax         DECIMAL(5, 2),
    prcp         DECIMAL(6, 2),
    wspd         DECIMAL(5, 2),
    pres         DECIMAL(7, 2),
    tsun         INTEGER,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (place_id, month)
);

CREATE INDEX idx_climate_normals_place_month ON climate_normals (place_id, month);

COMMENT
    ON COLUMN "climate_normals"."month" IS 'Range: 1 to 12';