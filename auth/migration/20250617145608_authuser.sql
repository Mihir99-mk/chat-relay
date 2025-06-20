-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_auth_users (
    id_uuid VARCHAR(40) PRIMARY KEY,
    provider_id INT NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,     -- e.g., Slack user ID, Google sub, GitHub ID
    team_id VARCHAR(255),                       -- For Slack or similar
    name VARCHAR(255),
    real_name VARCHAR(255),
    email VARCHAR(255),
    access_token TEXT NOT NULL,
    refresh_token TEXT,
    token_type VARCHAR(64),
    expires_at TIMESTAMP NULL,
    scope TEXT,
    raw_profile JSON,                           -- Store full original profile

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(100) NULL,
    updated_by VARCHAR(100) NULL,
    deleted_by VARCHAR(100) NULL,
    user_agent VARCHAR(255) NULL,
    ip_address VARCHAR(45) NULL,  

    CONSTRAINT fk_provider FOREIGN KEY (provider_id) REFERENCES tbl_auth_providers(id_int),
    UNIQUE KEY unique_provider_user (provider_id, provider_user_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
