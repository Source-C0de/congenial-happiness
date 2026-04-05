CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE departments (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(100) NOT NULL UNIQUE,
    color       VARCHAR(20)  NOT NULL DEFAULT '#1D9E75',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Seed default departments
INSERT INTO departments (name, color) VALUES
    ('Engineering', '#1D9E75'),
    ('Sales',       '#378ADD'),
    ('HR',          '#BA7517'),
    ('Finance',     '#D4537E'),
    ('IT',          '#7F77DD'),
    ('Operations',  '#888780');
