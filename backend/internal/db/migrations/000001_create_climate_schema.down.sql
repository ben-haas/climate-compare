-- create_climate_schema.down.sql
DROP INDEX IF EXISTS idx_climate_normals_place_month;
DROP TABLE IF EXISTS climate_normals;
DROP TABLE IF EXISTS places;