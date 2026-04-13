-- Migration: Add session_settings table for admin-configurable session policy
-- Admins can configure inactivity timeout and browser-close logout behavior.

CREATE TABLE IF NOT EXISTS session_settings (
    id              UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
    inactivity_timeout_minutes  INT         NOT NULL DEFAULT 30,   -- 0 = never expire
    logout_on_browser_close     BOOLEAN     NOT NULL DEFAULT FALSE,
    updated_by      UUID        REFERENCES users(id) ON DELETE SET NULL,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Insert default settings row (only one row ever exists)
INSERT INTO session_settings (inactivity_timeout_minutes, logout_on_browser_close)
VALUES (30, false)
ON CONFLICT DO NOTHING;
