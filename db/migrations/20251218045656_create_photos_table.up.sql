CREATE TABLE photos
(
    id         CHAR(36)     NOT NULL DEFAULT (UUID()) PRIMARY KEY,
    url        VARCHAR(255) NOT NULL,
    session_id CHAR(36)     NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_photos_sessions
        FOREIGN KEY (session_id)
            REFERENCES sessions (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;