
-- +migrate Up
CREATE TABLE comments(
    id bigserial not null,
    user_id bigint references users(id) not null,
    post_id bigint references posts(id) not null,
    content text not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    primary key(id)
);

-- +migrate Down
DROP TABLE comments;
