DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum_default') THEN
        CREATE TYPE status_enum_default AS ENUM (
            'active',
            'inactive'
        );
    END IF;
END
$$;

-- Table: account_tokens
-- for status 0 --> INACTIVE and 1 --> ACTIVE
CREATE TABLE IF NOT EXISTS account_tokens (
    id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    status status_enum_default NOT NULL DEFAULT 'inactive',
    expired_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX index_account_tokens ON account_tokens (username, id);
