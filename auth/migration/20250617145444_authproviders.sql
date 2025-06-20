-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_auth_providers (
    id_int INT NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id_int`),
    name VARCHAR(64) NOT NULL UNIQUE,         -- e.g., 'slack', 'google', 'github'
    display_name VARCHAR(128),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(100) NULL,
    updated_by VARCHAR(100) NULL,
    deleted_by VARCHAR(100) NULL,
    user_agent VARCHAR(255),
    ip_address VARCHAR(45)
);


INSERT INTO tbl_auth_providers (name, display_name) 
VALUES 
  ('slack', 'Slack'),
  ('google', 'Google'),
  ('github', 'GitHub');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
