CREATE TABLE sessions
(
    id           CHAR(36)     NOT NULL DEFAULT (UUID()) PRIMARY KEY,
    laundry_name VARCHAR(255) NOT NULL,
    status       VARCHAR(10)  NOT NULL,
    note         VARCHAR(255),
    dropped_at   TIMESTAMP    NULL     DEFAULT NULL,
    picked_up_at TIMESTAMP    NULL     DEFAULT NULL,
    user_id      CHAR(36)     NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_sessions_users
        FOREIGN KEY (user_id)
            REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;