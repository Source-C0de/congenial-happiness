CREATE TABLE employees (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name       VARCHAR(150)  NOT NULL,
    job_title       VARCHAR(150)  NOT NULL,
    department_id   UUID          NOT NULL REFERENCES departments(id) ON DELETE SET NULL,
    work_email      VARCHAR(255)  NOT NULL UNIQUE,
    extension       VARCHAR(10)   NOT NULL UNIQUE,
    mobile          VARCHAR(30),
    office_location VARCHAR(100),
    photo_url       VARCHAR(500),
    is_active       BOOLEAN       NOT NULL DEFAULT TRUE,
    online_status   VARCHAR(20)   NOT NULL DEFAULT 'offline'
                    CHECK (online_status IN ('online','away','offline')),
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_employees_department  ON employees(department_id);
CREATE INDEX idx_employees_email       ON employees(work_email);
CREATE INDEX idx_employees_extension   ON employees(extension);
CREATE INDEX idx_employees_active      ON employees(is_active);
-- Full text search index
CREATE INDEX idx_employees_search ON employees
    USING GIN (to_tsvector('english', full_name || ' ' || job_title || ' ' || work_email));
