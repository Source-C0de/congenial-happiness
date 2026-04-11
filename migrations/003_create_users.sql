-- "users" are employee accounts that can log into ContactHub
CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id     UUID          REFERENCES employees(id) ON DELETE SET NULL,
    email           VARCHAR(255)  NOT NULL UNIQUE,
    password_hash   VARCHAR(255)  NOT NULL,
    role            VARCHAR(20)   NOT NULL DEFAULT 'user'
                    CHECK (role IN ('admin', 'user')),
    is_active       BOOLEAN       NOT NULL DEFAULT TRUE,
    last_login_at   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role  ON users(role);

-- Seed admin account (password: Admin@1234 — change immediately in production)
INSERT INTO users (email, password_hash, role) VALUES
    ('admin@corp.com',
     '$2a$12$placeholder_replace_with_bcrypt_hash',
     'admin');
