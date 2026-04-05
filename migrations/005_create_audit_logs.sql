CREATE TABLE audit_logs (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    actor_user_id   UUID          REFERENCES users(id) ON DELETE SET NULL,
    actor_email     VARCHAR(255)  NOT NULL,
    action          VARCHAR(50)   NOT NULL,  -- created, updated, deleted, deactivated
    target_type     VARCHAR(50)   NOT NULL,  -- employee, user, department
    target_id       UUID,
    target_name     VARCHAR(255),
    metadata        JSONB,                   -- extra context as JSON
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_actor     ON audit_logs(actor_user_id);
CREATE INDEX idx_audit_action    ON audit_logs(action);
CREATE INDEX idx_audit_created   ON audit_logs(created_at DESC);

