CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE services (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL
);

CREATE TABLE reviews (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    rating SMALLINT NOT NULL,

    CONSTRAINT fk_reviews_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
