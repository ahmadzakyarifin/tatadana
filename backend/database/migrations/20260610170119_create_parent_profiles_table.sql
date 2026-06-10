-- +goose Up
-- +goose StatementBegin
CREATE TABLE parent_profiles (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    user_id INT UNSIGNED NOT NULL UNIQUE,

    nik VARCHAR(16) NULL UNIQUE,
    birth_place VARCHAR(50) NULL,
    birth_date DATE NULL,
    gender ENUM('male', 'female') NULL,

    address TEXT NULL,
    education VARCHAR(30) NULL,
    occupation VARCHAR(50) NULL,
    income DECIMAL(15,2) NULL,

    created_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,

    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    INDEX idx_parent_profiles_deleted_at (deleted_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS parent_profiles;
-- +goose StatementEnd
