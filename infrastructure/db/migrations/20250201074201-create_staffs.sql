
-- +migrate Up
CREATE TABLE staffs(
    id bigserial not null,
    name text not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    primary key(id)
);

-- +migrate Down
DROP TABLE staffs;
