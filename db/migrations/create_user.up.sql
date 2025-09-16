CREATE EXTENSION IF NOT EXISTS citext;
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
CREATE TYPE user_role AS ENUM ('user', 'admin');
END IF;
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned');
END IF;
END$$;

CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    username        CITEXT        NOT NULL,
    email           CITEXT        NOT NULL,
    password_hash   VARCHAR(255)  NOT NULL,
    role            user_role     NOT NULL DEFAULT 'user',
    status          user_status   NOT NULL DEFAULT 'active',
    full_name       VARCHAR(200),
    avatar_url      TEXT,
    last_login_at   TIMESTAMPTZ,
    last_login_ip   INET,
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ   NULL
    );

CREATE UNIQUE INDEX IF NOT EXISTS ux_users_email    ON users (email)    WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS ux_users_username ON users (username) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS ix_users_created_at ON users (created_at);
CREATE INDEX IF NOT EXISTS ix_users_status     ON users (status);
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_users_set_updated_at ON users;
CREATE TRIGGER trg_users_set_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();