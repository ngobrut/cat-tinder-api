CREATE TABLE
    IF NOT EXISTS users (
        user_id uuid default gen_random_uuid () not null constraint users_pk primary key,
        name varchar(255) not null,
        email varchar(255) not null,
        password varchar(500) not null,
        created_at timestamp default now (),
        updated_at timestamp default now (),
        deleted_at timestamp default null
    );

CREATE UNIQUE INDEX unique_email_idx ON users (email)
WHERE
    (deleted_at IS NULL);