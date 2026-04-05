CREATE TABLE sync_logs (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sync_type       VARCHAR(30)   NOT NULL CHECK (sync_type IN ('outlook','ucm')),
    status          VARCHAR(20)   NOT NULL CHECK (status IN ('success','failed','running')),
    records_affected INT          NOT NULL DEFAULT 0,
    error_message   TEXT,
    started_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    finished_at     TIMESTAMPTZ
);

