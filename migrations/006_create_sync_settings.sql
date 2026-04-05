CREATE TABLE sync_settings (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sync_type       VARCHAR(30)   NOT NULL UNIQUE CHECK (sync_type IN ('outlook','ucm')),
    is_enabled      BOOLEAN       NOT NULL DEFAULT FALSE,
    last_synced_at  TIMESTAMPTZ,
    config          JSONB
);

INSERT INTO sync_settings (sync_type, is_enabled) VALUES
    ('outlook', false),
    ('ucm',     false);
